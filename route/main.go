package route

import (
	"github.com/labstack/echo"
)


type Illuminate struct {
	echo.Echo
}

//type Route *echo.Group


//type Illuminate struct {
//	echo.Echo
//}

type Group struct {
	*echo.Group
}

type Resource struct {
	*Group
}

//func (route *Illuminate) Group(path string,handlerFunc echo.HandlerFunc,middlewareFunc ...echo.MiddlewareFunc) (*Group){
//	return *Group(path,handlerFunc,middlewareFunc)
//}

func (group *Group) Resource(txt string) (string) {
	return txt
}

func main() {


	//r := Illuminate{}
	//
	//r.Group("/").Resource("")


	//r := Illuminate{}
	//r2 := Group{}
	//r2.NewMethod()

	//r := Illuminate{}
	//r.Group("/").Resource()
	//
	//r2 := Illuminate{}
	//
	//
	//r1 := Route{}
	//
	//r1.Resource()
	//
	//
	//r.Resource()
}

type Route struct {
	*echo.Group
}
