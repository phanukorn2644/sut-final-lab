package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string // ต้องไม่เป็นค่าว่าง
	Email      string
	EmployeeID string // รหัสพนักงานขึ7นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว
}

func TestPositive(t *testing.T) {

	g := gomega.NewGomegaWithT(t)

	Employ := Employee{
		Name:       "Friend",
		Email:      "phanukorn@gmail.com",
		EmployeeID: "B6311308",
	}

	t.Run("Positive_test", func(t *testing.T) {

		ok, err := govalidator.ValidateStruct(Employ)
		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
	})
}
