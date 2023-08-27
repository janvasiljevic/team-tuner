package api

import (
	"jv/team-tone-tuner/router/middleware"
	"jv/team-tone-tuner/utils"

	"github.com/labstack/echo/v4"
)

func (api Api) RegisterRoutes(v1 *echo.Group) {

	protected := middleware.JWT(utils.JWTSecret)
	studentAuth := middleware.CheckUserRole(utils.UserRoleStudent)
	adminAuth := middleware.CheckUserRole(utils.UserRoleAdmin)

	githubLogin := v1.Group("/login/github")

	githubLogin.POST("", api.PostGithubLogin)
	githubLogin.GET("/callback", api.GetGithubLoginCallback)

	student := v1.Group("/student", protected, adminAuth)

	student.GET("", api.GetStudents)

	course := v1.Group("/course", protected, adminAuth)

	course.GET("", api.GetCourses)
	course.POST("", api.PostCourse)
	course.GET("/:courseId/stats/bf-box-plot", api.GetBigFiveBoxPlot)
	course.GET("/:courseId/stats/daily-activity", api.GetDailyActivity)
	course.GET("/:courseId/stats/questioner-stats", api.GetQuestionerCompletionStats)
	course.GET("/:courseId/group-runs", api.GetGroupRuns)

	group := v1.Group("/group", protected, adminAuth)

	group.GET("", api.GetGroups)
	group.GET("/:id", api.GetGroup)
	group.GET("/download/:groupRunId", api.GetDownloadGroupsCSV)
	group.POST("/form", api.PostFormGroups)

	questions := v1.Group("/question", protected, studentAuth)

	questions.GET("", api.GetQuestions)
	questions.POST("/submit", api.PostSubmitQuestions)
	questions.POST("/answer", api.PostQuestionsAnswer)

	auth := v1.Group("/auth", protected)

	auth.GET("/whoami", api.GetWhoAmI)
	auth.POST("/logout", api.Logout)

	answer := v1.Group("/answer", protected, studentAuth)

	answer.GET("", api.GetAllAnswers)
	answer.GET("/questioner-status", api.GetQuestionerStatus)
	answer.GET("/bfi-report", api.GetBfiReport)
}
