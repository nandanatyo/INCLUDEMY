package rest

import (
	"fmt"
	"includemy/internal/service"
	"includemy/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoints() {
	r.router.Use(r.middleware.Timeout())
	user := r.router.Group("/user", r.middleware.AuthenticateUser)
	profile := user.Group("/profile", r.middleware.AuthenticateUser)
	admin := r.router.Group("/admin", r.middleware.AuthenticateUser, r.middleware.OnlyAdmin)
	search := r.router.Group("/search")

	//User
	r.router.POST("/signin", r.Signin)           //register
	r.router.POST("/login", r.Login)             //login
	profile.POST("/upload-photo", r.UploadPhoto) //upload photo profile user
	profile.POST("/update-user", r.UpdateUser)   //mengupdate profile user

	admin.DELETE("/delete-user/:id", r.DeleteUser) //menghapus user

	//Course-Subcourse
	search.GET("/course/", r.GetCourseByAny) //melihat Course berdasarkan id atau title

	user.GET("/course/subcourse", r.GetSubCourseWithinCourse)  //melihat subCourse dalam Course
	user.POST("/join-course", r.CreateUserJoinCourse)          //mendaftar Course
	user.POST("/join-course/subcourse", r.CreateUserSubcourse) //user otomatis mendaftar subCourse (logika dari FE)
	user.GET("/course", r.GetUserCourse)                       //mendapatkan course yang user ikuti
	user.PATCH("/update-subcourse/:id", r.UpdateUserSubcourse) //mengupdate subCourse dari user apakah di checklist atau tidak

	admin.POST("/create-course", r.CreateCourse)                                     //membuat Course
	admin.PATCH("/update-course/:id", r.UpdateCourse)                                //mengupdate Course
	admin.POST("/create-course/upload-file", r.UploadCoursePhoto)                    //upload file Course
	admin.DELETE("/delete-course/:id", r.DeleteCourse)                               //menghapus Course
	admin.POST("/create-course/create-subcourse", r.CreateSubcourse)                 //membuat subCourse
	admin.PATCH("/update-subcourse/:id", r.UpdateSubcourse)                          //mengupdate subCourse
	admin.POST("/create-course/create-subcourse/upload-file", r.UploadSubcourseFile) //upload file subCourse
	admin.DELETE("/delete-subcourse/:id", r.DeleteSubcourse)                         //menghapus subCourse
	admin.DELETE("/delete-user-join-course/:id", r.DeleteUserJoinCourse)             //menghapus user join Course
	admin.GET("/user/subkursus-on-one-course/:id", r.GetUserSubCourseOnOneCourse)    //melihat subCourse yang dimiliki user, multiple return

	//Sertification
	search.GET("/sertification/", r.GetSertificationByTitleOrID) //melihat sertification berdasarkan id atau title

	user.POST("/create-sertification-user", r.CreatSertificationUser) //user registrasi ke sertification
	user.GET("/sertification", r.GetUserSertification)                //mendapatkan sertification yang diregistrasi user

	admin.POST("/create-sertification", r.CreateSertification)                //membuat sertification
	admin.DELETE("/delete-sertification/:id", r.DeleteSertification)          //menghapus sertification
	admin.PATCH("/update-sertification/:id", r.UpdateSertification)           //mengupdate sertification
	admin.POST("/create-sertification/upload-file", r.UploadSertifPhoto)      //upload file sertification
	admin.DELETE("/delete-sertification-user/:id", r.DeleteSertificationUser) //menghapus sertification user

	//ApplyJob
	search.GET("/job/", r.GetJobByTitleOrID) //melihat job berdasarkan id atau title

	user.POST("/apply-job", r.CreateApplicant)     //user melamar job
	user.GET("/application", r.GetUserApplication) //melihat application yang dilakukan user

	admin.POST("/create-job", r.CreateJob)                 //membuat job
	admin.DELETE("/delete-job/:id", r.DeleteJob)           //menghapus job
	admin.PATCH("/update-job/:id", r.UpdateJob)            //mengupdate job
	admin.POST("/create-job/upload-file", r.UploadJobFile) //upload file job

}

func (r *Rest) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	r.router.Run(fmt.Sprintf(":%s", port))
}
