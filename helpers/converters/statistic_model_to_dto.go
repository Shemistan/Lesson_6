package converters

import (
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/models"
)

func StatisticModelToDto(model models.Statistic) dtos.Statistics {
	return dtos.Statistics{
		DeletedUsersCount: model.DeletedUsersCount,
		UpdateCount:       model.UpdateCount,
		GetUserCount:      model.GetUserCount,
		GetUsersCount:     model.GetUsersCount,
	}
}
