package route

import (
	"github.com/kataras/iris"
	"golang123/config"
	"golang123/controller/common"
	"golang123/controller/auth"
	"golang123/controller/category"
	"golang123/controller/article"
	"golang123/controller/collect"
	"golang123/controller/comment"
	"golang123/controller/vote"
	"golang123/controller/user"
	"golang123/controller/message"
)

// Route 路由
func Route(app *iris.Application) {
	apiPrefix   := config.APIConfig.Prefix

	routes := app.Party(apiPrefix) 
	{	
		routes.Post("/signin",                   user.Signin)
		routes.Post("/signup",                   user.Signup)
		routes.Post("/signout",                  user.Signout)
		routes.Get("/active/verify/:id/:secret", user.VerifyActiveLink)
		routes.Post("/active/:id/:secret",       user.ActiveAccount)
		routes.Post("/reset",                    user.ResetPasswordMail)
		routes.Get("/reset/verify/:id/:secret",  user.VerifyResetPasswordLink)
		routes.Post("/reset/:id/:secret",        user.ResetPassword)
		routes.Get("/heartbeat",                 common.Heartbeat)

		routes.Get("/user/info/public/:id",  user.PublicInfo)
		routes.Get("/user/info",             auth.SigninRequired,  
											 user.Info)
		routes.Post("/user/update",          auth.ActiveRequired,       
										     user.UpdateInfo)
		routes.Post("/user/password/update", auth.ActiveRequired,       
											 user.UpdatePassword)
		routes.Get("/user/score/top10",      user.Top10)
		routes.Get("/user/score/top100",     user.Top100)

		routes.Post("/upload",               auth.ActiveRequired,          
											 common.Upload)

		routes.Get("/message/unread",        auth.SigninRequired,  
											 message.Unread)
		routes.Get("/message/unread/count",  auth.SigninRequired,  
											 message.UnreadCount)

		routes.Get("/categories",            category.List)

		routes.Get("/articles",                article.List)
		routes.Get("/articles/user/:userID",   article.UserArticleList)
		routes.Get("/articles/maxcomment",     article.ListMaxComment)
		routes.Get("/articles/maxbrowse",      article.ListMaxBrowse)
		routes.Get("/article/{id:int min(1)}", article.Info)
		routes.Post("/article/create",         auth.ActiveRequired, 
										       article.Create)
		routes.Post("/article/update",         auth.ActiveRequired,    
											   article.Update)
		routes.Post("/article/delete/:id",     auth.ActiveRequired,
											   article.Delete)
		routes.Get("/articles/top",            article.Tops)
		routes.Post("/article/top/:id",        auth.EditorRequired,    
											   article.Top)
		routes.Post("/article/deltop/:id",     auth.EditorRequired,    
											   article.DeleteTop)
											   
		routes.Post("/collect/create",      auth.ActiveRequired,
											collect.Collect)
		routes.Post("/collect/delete",      auth.ActiveRequired,
										    collect.DeleteCollect)
		routes.Get("/collects",             auth.SigninRequired,
											collect.List)

		routes.Post("/comment/create",       auth.ActiveRequired,
											 comment.Create)
		routes.Get("/comments/user/:userID", comment.UserCommentList)

		routes.Get("/votes",                vote.List)
		routes.Get("/votes/maxbrowse",      vote.ListMaxBrowse)
		routes.Get("/votes/maxcomment",     vote.ListMaxComment)
		routes.Get("/votes/user/:userID",   vote.UserVoteList)
		routes.Post("/vote/create",         auth.EditorRequired,
											vote.Create)
		routes.Post("/vote/delete",         auth.EditorRequired,
											vote.Delete)
		routes.Get("/vote/:id",             vote.Info)
		routes.Post("/vote/item/create",    auth.EditorRequired,
											vote.CreateVoteItem)
		routes.Post("/vote/item/edit",      auth.EditorRequired,
											vote.EditVoteItem)
		routes.Post("/vote/uservote/:id",   auth.ActiveRequired,
											vote.UserVoteVoteItem)
    }

	adminRoutes := app.Party(apiPrefix + "/admin", auth.AdminRequired)
	{
		adminRoutes.Get("/categories",               category.AllList)
		adminRoutes.Post("/category/create",         category.Create)
		adminRoutes.Post("/category/update",         category.Update)
		adminRoutes.Post("/category/status/update",  category.UpdateStatus)

		adminRoutes.Get("/articles",                 article.AllList)
		adminRoutes.Post("/article/status/update",   article.UpdateStatus)
    }
}