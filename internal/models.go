package internal

type GTFS struct {
	agencies      []Agency
	routes        []Route
	calendars     []Calendar
	calendarDates []CalendarDate
	shapes        []Shape
}

type Agency struct {
	ID       string `csv:"agency_id" db:"id"`
	Name     string `csv:"agency_name" db:"name"`
	URL      string `csv:"agency_url" db:"url"`
	Timezone string `csv:"agency_timezone" db:"timezone"`
	Language string `csv:"agency_lang" db:"language"`
	Phone    string `csv:"agency_phone" db:"phone"`
}

type Route struct {
	ID        string `csv:"route_id" db:"id"`
	AgencyID  string `csv:"agency_id" db:"agency_id"`
	ShortName string `csv:"route_short_name" db:"short_name"`
	LongName  string `csv:"route_long_name" db:"long_name"`
	Type      int    `csv:"route_type" db:"type"`
	Color     string `csv:"route_color" db:"color"`
	TextColor string `csv:"route_text_color" db:"text_color"`
}

type Calendar struct {
	ServiceID string `csv:"service_id" db:"service_id"`
	StartDate string `csv:"start_date" db:"start_date"`
	EndDate   string `csv:"end_date" db:"end_date"`
	Monday    string `csv:"monday" db:"monday"`
	Tuesday   int    `csv:"tuesday" db:"tuesday"`
	Wednesday int    `csv:"wednesday" db:"wednesday"`
	Thursday  int    `csv:"thursday" db:"thursday"`
	Friday    int    `csv:"friday" db:"friday"`
	Saturday  int    `csv:"saturday" db:"saturday"`
	Sunday    int    `csv:"sunday" db:"sunday"`
}

type CalendarDate struct {
	ServiceID     string `csv:"service_id" db:"service_id"`
	Date          string `csv:"date" db:"date"`
	ExceptionType string `csv:"exception_type" db:"exception_type"`
}

type Shape struct {
	ShapeID          string  `csv:"shape_id" db:"shape_id"`
	Latitude         string  `csv:"shape_pt_lat" db:"latitude"`
	Longitude        string  `csv:"shape_pt_lon" db:"longitude"`
	Sequence         int     `csv:"shape_pt_sequence" db:"sequence"`
	DistanceTraveled float64 `csv:"shape_dist_traveled" db:"distance_traveled"`
}

type StopTime struct {
	TripID string `csv:"trip_id" db:"trip_id"`
	ArrivalTime string `csv:"arrival_time" db:"arrival_time"`
	ArrivalTime string `csv:"arrival_time" db:"arrival_time"`
}

func (r Route) TypeEmoji() string {
	if r.Type == 0 {
		return "🚋 (tram)"
	}

	if r.Type == 1 {
		return "🚇 (metro)"
	}

	if r.Type == 2 {
		return "🚃 (rail)"
	}

	if r.Type == 3 {
		return "🚌 (bus)"
	}

	if r.Type == 4 {
		return "⛴ (ferry)"
	}

	if r.Type == 5 {
		return "🚎 (cable car)"
	}

	if r.Type == 6 {
		return "🚠 (gondola)"
	}

	if r.Type == 7 {
		return "🚈 (Funicular)"
	}

	return "🚙"
}
