package controllers

import (
	"github.com/gorilla/context"
	"html/template"
	"net/http"
)

func (c *Controller) MyPageGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := context.Get(r, "id").(int)
		pet, err := c.UserStore.GetPet(userID) // you can get pet as first param of this method
		if err != nil {
			http.Redirect(w, r, "/petcabinet", http.StatusFound)
			return
		}
			


			tmpl, _ := template.ParseFiles("web/templates/mypage.html")
			tmpl.Execute(w, pet)
			blog:=c.BlogStore.GetBlog()
			tmp2,_:=template.ParseFiles("web/templates/blog.html")
			tmp2.Execute(w, blog)

		//view.GenerateHTML(w, nil, "mypage")
	}
}


