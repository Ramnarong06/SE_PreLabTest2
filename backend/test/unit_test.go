package unit

import (
	"testing"
	
	"github.com/Ramnarong06/SE_PreLabTest2/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	

)

func TestUnitUser(t *testing.T) {
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

func TestUnitEmailUser(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run("Email error", func(t *testing.T){
		user := entity.User{
			Username:  	"Ramnarong",
			Email:	  	"b6513214", // ผิดตรงนี้
			Password:	"123456",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid")) // ข้อความที่กำหนดใน Entity
	})
}

func TestUnitUserName(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run("Username error", func(t *testing.T){
		user := entity.User{
			Username:  	"", // ผิดตรงนี้
			Email:	  	"b6513214@gmail.com",
			Password:	"123456",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Username is required"))
	})
}