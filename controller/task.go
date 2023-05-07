package controller

import (
	"net/http"
	"mini_project/database"
	"mini_project/middleware"
	"mini_project/model"
	"mini_project/util"

	"github.com/labstack/echo/v4"
)

// Nambah Task
func CreateTaskController(c echo.Context) error {
	task := model.Task{}
	c.Bind(&task)

	errs := util.TaskValidate(task)
	if errs != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Failed create task input invalid",
			"Errors":  errs,
		})
	}

	UserID := middleware.ExtractTokenUserId(c)

	task.UserID = uint(UserID)

	if err := database.DB.Save(&task).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new task",
		"task":    task,
	})
}

// Melihat Tugas
func GetTasksController(c echo.Context) error {
	var tasks []model.Task

	UserID := middleware.ExtractTokenUserId(c)

	if err := database.DB.Where("user_id = ?", UserID).Preload("User").Find(&tasks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all tasks",
		"tasks":   tasks,
	})
}

// Update Tugas
func UpdateTaskController(c echo.Context) error {
	task := model.Task{}
	TaskID, err := util.ConvertToInt(c.Param("id"))
	if err != nil {
		return err
	}

	if err := database.DB.First(&task, TaskID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Bind(&task)
	if err := database.DB.Save(&task).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "task updated successfully",
	})
}

// Menghapus Tugas
func DeleteTaskController(c echo.Context) error {
	task := model.Task{}
	TaskId, err := util.ConvertToInt(c.Param("id"))
	if err != nil {
		return err
	}

	if err := database.DB.Delete(&task, TaskId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Task deleted successfully",
	})
}

// Search Tugas Berdasarkan Nama
func GetTaskController(c echo.Context) error {
	var tasks []model.Task
	name := c.QueryParam("name")

	UserID := middleware.ExtractTokenUserId(c)

	if err := database.DB.Where("user_id = ? AND name LIKE ?", UserID, "%"+name+"%").Find(&tasks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get task by name",
		"task":    tasks,
	})
}

// Menampilkan Timeline
func ShowTimelineController(c echo.Context) error {
	var tasks []model.Task

	deadlineStr := c.QueryParam("deadline")
	deadlineDays, err := util.ConvertToInt(deadlineStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid deadline"})
	}

	TimeStart, TimeFinish := util.DeadlineRange(deadlineDays)
	UserID := middleware.ExtractTokenUserId(c)

	if err := database.DB.Where("user_id = ? AND status = ? AND due_date BETWEEN ? AND ?", UserID, "unfinished", TimeStart, TimeFinish).Order("due_date").Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to query tasks"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all tasks",
		"tasks":   tasks,
	})
}

// Change Status Tugas by ID
func ChangeStatusController(c echo.Context) error {
	task := model.Task{}
	TaskId, err := util.ConvertToInt(c.Param("id"))
	if err != nil {
		return err
	}

	if err := database.DB.Model(&task).Where("id = ?", TaskId).Update("status", "finished").Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Task status changed successfully",
	})
}
