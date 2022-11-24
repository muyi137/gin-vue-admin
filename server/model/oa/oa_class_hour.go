// 自动生成模板OaClassHour
package oa

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "time"
)

// OaClassHour 结构体
type OaClassHour struct {
    global.GVA_MODEL
    Card_number string     `json:"card_number" form:"card_number" gorm:"column:card_number;comment:身份证号;size:20;"`
    ClassTime   *int       `json:"classTime" form:"classTime" gorm:"column:class_time;comment:课头(几门课);size:6;"`
    ReportTime  *time.Time `json:"reportTime" form:"reportTime" gorm:"column:report_time;comment:上报日期;"`
    Score       *float64   `json:"score" form:"score" gorm:"column:score;comment:;size:6;"`
    Score2      *float64   `json:"score2" form:"score2" gorm:"column:score2;comment:k1体;size:6;"`
    Score1      *float64   `json:"score1" form:"score1" gorm:"column:score1;comment:k1理;size:6;"`
    Score3      *float64   `json:"score3" form:"score3" gorm:"column:score3;comment:k1机;size:6;"`
    Score4      *float64   `json:"score4" form:"score4" gorm:"column:score4;comment:k2理;size:6;"`
    Score5      *float64   `json:"score5" form:"score5" gorm:"column:score5;comment:k2机;size:6;"`
    Score6      *float64   `json:"score6" form:"score6" gorm:"column:score6;comment:k3理;size:6;"`
    Score7      *float64   `json:"score7" form:"score7" gorm:"column:score7;comment:k3机;size:6;"`
    Score8      *float64   `json:"score8" form:"score8" gorm:"column:score8;comment:竞赛折算;size:6;"`
    Score9      *float64   `json:"score9" form:"score9" gorm:"column:score9;comment:自定义1;size:6;"`
    Score10     *float64   `json:"score10" form:"score10" gorm:"column:score10;comment:自定义2;size:6;"`
    Score11     *float64   `json:"score11" form:"score11" gorm:"column:score11;comment:自定义3;size:6;"`
    Score12     *float64   `json:"score12" form:"score12" gorm:"column:score12;comment:自定义4;size:6;"`
    Status      string     `json:"status" form:"status" gorm:"column:status;type:enum('已审核','待审核');comment:状态;"`
    CreatedBy   uint       `gorm:"column:created_by;comment:创建者"`
    UpdatedBy   uint       `gorm:"column:updated_by;comment:更新者"`
    DeletedBy   uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName OaClassHour 表名
func (OaClassHour) TableName() string {
    return "oa_class_hour"
}
