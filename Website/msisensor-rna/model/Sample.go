package model

import (
	"msisensor-rna/utils/errmsg"

	"gorm.io/gorm"
)

type Sample struct {
	User User `gorm:"foreignkey:Uid"`
	gorm.Model
	Uid         int     `gorm:"type:int;not null" json:"uid"`
	PatientName string  `gorm:"type:varchar(20);not null" json:"patientname"`
	FileName    string  `gorm:"type:varchar(20)" json:"filename"`
	CancerType  string  `gorm:"type:varchar(20)" json:"cancertype"`
	Gender      int     `gorm:"type:int" json:"gender"`
	Age         int     `gorm:"type:int;DEFAULT:0" json:"age"`
	Prediction  int     `gorm:"type:int;DEFAULT:2" json:"prediction"`
	Weight      string  `gorm:"type:text" json:"weight"`
	Value       string  `gorm:"type:text" json:"value"`
	Score       float32 `gorm:"type:float;DEFAULT:-1.0" json:"score"`
}
type List struct {
}

// 新增样本
func CreateSample(data *Sample) int {
	err := db.Create(&data).Error
	// weight, err := json.Marshal(data.Weight)
	// if err != nil {
	// 	fmt.Println("json marshal error:", err)
	// }
	// data.Weight = weight

	// if err != nil {
	// 	fmt.Println("json marshal error:", err)
	// }
	// data.Value = data.Value
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// 查询该用户的样本列表
func GetSamples(uid int) ([]Sample, int) {
	var sampleList []Sample
	// var total int64

	// err = db.Where("uid = ?", uid).Select("sample.id, created_at, updated_at, samp_name, cancer_type, samp_pre").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&sampleList).Error
	err = db.Where("uid = ?", uid).Select("sample.id, gender, age, created_at, updated_at, cancer_type, prediction, patient_name,file_name,weight,value,score").Order("Created_At DESC").Find(&sampleList).Error

	// 单独计数
	// db.Model(&sampleList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR
	}
	return sampleList, errmsg.SUCCESS
}

// 查询样本
func GetSample(sid int) (Sample, int) {
	var sample Sample
	err := db.Where("id = ?", sid).Preload("User").Find(&sample).Error
	if err != nil {
		return sample, errmsg.ERROR_SAMP_NOT_EXIST
	}
	return sample, errmsg.SUCCESS
}

// 查询样本是否存在
func CheckSample(sampName string) (code int) {
	var sample Sample
	db.Select("id").Where("sampName = ?", sampName).First(&sample)
	if sample.ID > 0 {
		return errmsg.ERROR_SAMPNAME_USED
	}
	return errmsg.SUCCESS
}

// 删除样本
func DeleteSample(sid int) int {
	var sample Sample
	err = db.Where("id = ? ", sid).Delete(&sample).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑样本
func EditSample(sid int, data *Sample) int {
	var sample Sample
	var maps = make(map[string]interface{})
	if data.FileName != "" {
		maps["FileName"] = data.FileName
	}
	if data.PatientName != "" {
		maps["PatientName"] = data.PatientName
	}
	if data.CancerType != "" {
		maps["CancerType"] = data.CancerType
	}
	if data.Gender != 0 {
		maps["Gender"] = data.Gender
	}
	if data.Age != 0 {
		maps["Age"] = data.Age
	}
	if data.Prediction != 2 {
		maps["Prediction"] = data.Prediction
	}
	if data.Weight != "" {
		maps["Weight"] = data.Weight
	}
	if data.Value != "" {
		maps["Value"] = data.Value
	}
	if data.Score != -1 {
		maps["Score"] = data.Score
	}

	err = db.Model(&sample).Where("id = ? ", sid).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
