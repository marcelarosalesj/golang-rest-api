package main

import (
	"os"

	"github.com/marcelarosalesj/golang-rest-api/controller"
	router "github.com/marcelarosalesj/golang-rest-api/http"
	"github.com/marcelarosalesj/golang-rest-api/repository"
	"github.com/marcelarosalesj/golang-rest-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(os.Getenv("PORT"))
}
