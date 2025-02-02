package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/ZupIT/horusec-platform/analytic/cmd/migration/v2/enums"
	dashboardController "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
	dashboardRepository "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	dashboardUseCases "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

type AnalyticMigration struct {
	dbConnectionAnalytic *database.Connection
	dbConnectionHorusec  *database.Connection
	dashboardController  dashboardController.IController
	summary              map[string][]string
}

func NewAnalyticMigrationV2() *AnalyticMigration {
	analyticMigration := &AnalyticMigration{
		dbConnectionAnalytic: getAnalyticDatabaseConnection(),
		dbConnectionHorusec:  getHorusecDatabaseConnection(),
		summary:              setSummary(),
	}

	analyticMigration.dashboardController = dashboardController.NewDashboardController(
		dashboardRepository.NewRepoDashboard(analyticMigration.dbConnectionAnalytic),
		analyticMigration.dbConnectionAnalytic, dashboardUseCases.NewUseCaseDashboard())

	return analyticMigration
}

func setSummary() map[string][]string {
	return map[string][]string{
		enums.SummarySuccess: {},
		enums.SummaryFailed:  {},
	}
}

func getAnalyticDatabaseConnection() *database.Connection {
	analyticDatabaseConfig := databaseConfig.NewDatabaseConfig()
	analyticDatabaseConfig.SetURI(
		env.GetEnvOrDefault(enums.EnvAnalyticDatabaseURI, enums.DefaultAnalyticDatabaseURIValue))

	dbConnectionAnalytic, err := database.NewDatabaseReadAndWrite(analyticDatabaseConfig)
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToAnalyticDatabase, err)
	}

	return dbConnectionAnalytic
}

func getHorusecDatabaseConnection() *database.Connection {
	horusecDatabaseConfig := databaseConfig.NewDatabaseConfig()
	horusecDatabaseConfig.SetURI(env.GetEnvOrDefault(enums.EnvHorusecDatabaseURI, enums.DefaultHorusecDatabaseURIValue))

	dbConnectionHorusec, err := database.NewDatabaseReadAndWrite(horusecDatabaseConfig)
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToHorusecDatabase, err)
	}

	return dbConnectionHorusec
}

func (a *AnalyticMigration) getAllAnalysisWithVulnerabilities() []*analysisEntities.Analysis {
	analysis := &[]*analysisEntities.Analysis{}

	preloads := map[string][]interface{}{
		"AnalysisVulnerabilities":               {},
		"AnalysisVulnerabilities.Vulnerability": {},
	}

	if err := a.dbConnectionHorusec.Read.FindPreload(analysis, map[string]interface{}{},
		preloads, "analysis").GetError(); err != nil {
		logger.LogPanic(enums.MessageFailedToGetAllAnalysis, err)
	}

	return *analysis
}

func (a *AnalyticMigration) loggingRegisterBeingMigrated(analysis *analysisEntities.Analysis) {
	message := fmt.Sprintf(enums.MessageRegisterBeingMigrated, analysis.WorkspaceID, analysis.WorkspaceName,
		analysis.RepositoryID, analysis.RepositoryName, analysis.CreatedAt, len(analysis.AnalysisVulnerabilities))

	logger.LogInfo(message)
}

func (a *AnalyticMigration) setMigrationInSummary(analysisID uuid.UUID, err error, table string) {
	if err == nil {
		message := fmt.Sprintf("analysis with id '%s' migrated with success on table '%s'",
			analysisID.String(), table)
		a.summary[enums.SummarySuccess] = append(a.summary[enums.SummarySuccess], message)
		return
	}

	message := fmt.Sprintf("failed to migrate analsysis with id '%s' on table '%s' with error -> '%v'",
		analysisID.String(), table, err)
	a.summary[enums.SummaryFailed] = append(a.summary[enums.SummaryFailed], message)
}

func (a *AnalyticMigration) migrateAnalysis(analysis *analysisEntities.Analysis) {
	a.setMigrationInSummary(analysis.ID, a.dashboardController.AddVulnerabilitiesByAuthor(analysis),
		dashboardEnums.TableVulnerabilitiesByAuthor)

	a.setMigrationInSummary(analysis.ID, a.dashboardController.AddVulnerabilitiesByLanguage(analysis),
		dashboardEnums.TableVulnerabilitiesByLanguage)

	a.setMigrationInSummary(analysis.ID, a.dashboardController.AddVulnerabilitiesByRepository(analysis),
		dashboardEnums.TableVulnerabilitiesByRepository)

	a.setMigrationInSummary(analysis.ID, a.dashboardController.AddVulnerabilitiesByTime(analysis),
		dashboardEnums.TableVulnerabilitiesByTime)
}

func (a *AnalyticMigration) printResults(total int) {
	a.createResultLog()

	fmt.Println()
	logger.LogWarn("MIGRATION FINISHED! CHECK THE RESULTS -->")
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS TO MIGRATE: %d", total))
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS SUCCESSFULLY MIGRATED: %d", len(a.summary[enums.SummarySuccess])))
	logger.LogWarn(fmt.Sprintf("TOTAL RECORDS THAT FAILED TO MIGRATE: %d", len(a.summary[enums.SummaryFailed])))
	logger.LogWarn("YOU CAN SEE THE COMPLETE RESULT IN '/tmp/v1-to-v2-horusec-analytic-result'")
}

func (a *AnalyticMigration) createResultLog() {
	result := "RESULT HORUSEC ANALYTIC MIGRATION V1 TO V2\n\nANALYSIS ID MIGRATED WITHOUT ERRORS -->\n"
	for _, value := range a.summary[enums.SummarySuccess] {
		result += fmt.Sprintf("SUCCESS: %s\n", value)
	}

	result += "\nANALYSIS ID AND TABLE THAT FAILED TO MIGRATE -->\n"
	for _, value := range a.summary[enums.SummaryFailed] {
		result += fmt.Sprintf("FAILED: %s\n", value)
	}

	file, _ := os.Create("./tmp/v1-to-v2-horusec-analytic-result")
	_, _ = file.WriteString(result)
}

func main() {
	analyticMigration := NewAnalyticMigrationV2()

	allPastAnalysis := analyticMigration.getAllAnalysisWithVulnerabilities()
	for _, analysis := range allPastAnalysis {
		analyticMigration.loggingRegisterBeingMigrated(analysis)
		analyticMigration.migrateAnalysis(analysis)
	}

	analyticMigration.printResults(len(allPastAnalysis) * enums.TotalOfTables)
}
