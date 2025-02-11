package models

import (
	"github.com/orotig/server_demo/config"
)

func GetWelderOrders(welder_orders *[]WelderTest) (err error) {
	if err = config.PGDB.Find(welder_orders).Error; err != nil {
		return err
	}
	return nil
}

func GetPendingWelderOrders(welder_orders *[]WelderTest) (err error) {
	if err = config.PGDB.Where("is_complete = ?", "false").Find(welder_orders).Error; err != nil {
		return err
	}
	return nil
}

func UpdateWelderOrder(welder_order *WelderTest) (err error) {
	if err = config.PGDB.Save(welder_order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrders(marker_orders *[]MarkerTest) (err error) {
	if err = config.PGDB.Find(marker_orders).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOrder(marker_order *MarkerTest) (err error) {

	if err = config.PGDB.Save(marker_order).Error; err != nil {
		return err
	}
	return nil
}
