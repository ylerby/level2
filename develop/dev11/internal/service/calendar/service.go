package calendar

import (
	"fmt"
	"strconv"
)

func (c *Calendar) Create(event Events, date string) error {
	if _, ok := c.CalendarMap[date]; ok {
		for _, value := range c.CalendarMap[date] {
			if value.UserId == event.UserId && value.EventName == event.EventName {
				return fmt.Errorf("element already exists")
			}
		}
	}
	c.CalendarMap[date] = append(c.CalendarMap[date], event)
	return nil
}

func (c *Calendar) Update(event Events, date string) error {
	if _, ok := c.CalendarMap[date]; !ok {
		return fmt.Errorf("element not found")
	} else {
		flag := false
		for index, value := range c.CalendarMap[date] {
			if value.UserId == event.UserId || value.EventName == event.EventName {
				flag = true
				c.CalendarMap[date][index] = event
			}
		}
		if !flag {
			return fmt.Errorf("element not found")
		}
	}
	return nil
}

func (c *Calendar) Delete(event Events, date string) error {
	_, ok := c.CalendarMap[date]
	if ok {
		for index, value := range c.CalendarMap[date] {
			if value.UserId == event.UserId && value.EventName == event.EventName {
				c.CalendarMap[date] = append(c.CalendarMap[date][:index], c.CalendarMap[date][index+1:]...)
				if len(c.CalendarMap[date]) == 0 {
					delete(c.CalendarMap, date)
				}
			}
		}
	} else {
		return fmt.Errorf("element not found")
	}
	return nil
}

func (c *Calendar) GetEventsForDay(day, month, year string) ([]Events, error) {
	if value, ok := c.CalendarMap[fmt.Sprintf("%s-%s-%s", day, month, year)]; !ok {
		return nil, fmt.Errorf("element not found")
	} else {
		return value, nil
	}
}

func (c *Calendar) GetEventsForWeek(day, month, year string) ([][]Events, error) {
	result := make([][]Events, 0)

	dayNumber, err := strconv.Atoi(day)
	if err != nil {
		return nil, fmt.Errorf("invalid data")
	}

	flag := false
	for i := 0; i < 7; i++ {
		res := strconv.Itoa(dayNumber + i)

		if _, ok := c.CalendarMap[fmt.Sprintf("%s-%s-%s", res, month, year)]; ok {
			result = append(result, c.CalendarMap[fmt.Sprintf("%s-%s-%s", res, month, year)])
			flag = true
		}
	}

	if !flag || len(result) == 0 {
		return nil, fmt.Errorf("element not found")
	}

	return result, nil
}

func (c *Calendar) GetEventsForMonth(month, year string) ([][]Events, error) {
	result := make([][]Events, 0)
	flag := false
	for i := 1; i < 32; i++ {
		res := strconv.Itoa(i)
		if _, ok := c.CalendarMap[fmt.Sprintf("%s-%s-%s", res, month, year)]; ok {
			result = append(result, c.CalendarMap[fmt.Sprintf("%d-%s-%s", i, month, year)])
			flag = true
		}
	}

	if !flag || len(result) == 0 {
		return nil, fmt.Errorf("element not found")
	}

	return result, nil
}
