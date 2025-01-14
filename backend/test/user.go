package uint

import (
	"testing"
	
	"github.com/Ramnarong06/SE_PreLabTest2/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	

)

func TestUintUser(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Correct", func(t *testing.T){
		user := entity.User{
			Username:  	"Ramnarong",
			Email:	  	"b6513214@gmail.com",
			Password:	"123456",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}