package calendar

import "fmt"

func (c *Calendar) Create(currentEventName string, event Events) error {
	if _, ok := c.CalendarMap[currentEventName]; ok {
		return fmt.Errorf("element alredy exist")
	}
	c.CalendarMap[currentEventName] = event
	return nil
}

func (c *Calendar) Update(currentEventName string, event Events) error {
	if _, ok := c.CalendarMap[currentEventName]; !ok {
		flag := false
		for _, value := range c.CalendarMap {
			if value.UserId == event.UserId || value.Date == event.Date {
				flag = true
				c.CalendarMap[currentEventName] = event
			}
		}
		if !flag {
			return fmt.Errorf("element not found")
		}
	} else {
		c.CalendarMap[currentEventName] = event
	}
	return nil
}

func (c *Calendar) Delete(currentEventName string, event Events) error {
	_, ok := c.CalendarMap[currentEventName]
	if ok {
		delete(c.CalendarMap, currentEventName)
	} else {
		flag := false
		for key, value := range c.CalendarMap {
			if value.UserId == event.UserId && value.Date == event.Date {
				flag = true
				delete(c.CalendarMap, key)
			}
		}

		if !flag {
			return fmt.Errorf("element not found")
		}
	}
	return nil
}

func (c *Calendar) GetEventsForDay(day string)     {}
func (c *Calendar) GetEventsForWeek(week string)   {}
func (c *Calendar) GetEventsForMonth(month string) {}
