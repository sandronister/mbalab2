package dto

type ValueDTO struct {
	Value string `json:"value"`
}

type CurrentConditionDTO struct {
	FeelsLikeC       string     `json:"FeelsLikeC"`
	FeelsLikeF       string     `json:"FeelsLikeF"`
	Cloudcover       string     `json:"cloudcover"`
	Humidity         string     `json:"humidity"`
	LangPt           []ValueDTO `json:"lang_pt"`
	LocalObsDateTime string     `json:"localObsDateTime"`
	ObservationTime  string     `json:"observation_time"`
	PrecipInches     string     `json:"precipInches"`
	PrecipMM         string     `json:"precipMM"`
	Pressure         string     `json:"pressure"`
	PressureInches   string     `json:"pressureInches"`
	TempC            string     `json:"temp_C"`
	TempF            string     `json:"temp_F"`
	UvIndex          string     `json:"uvIndex"`
	Visibility       string     `json:"visibility"`
	VisibilityMiles  string     `json:"visibilityMiles"`
	WeatherCode      string     `json:"weatherCode"`
	WeatherDesc      []ValueDTO `json:"weatherDesc"`
	WeatherIconURL   []ValueDTO `json:"weatherIconUrl"`
	Winddir16Point   string     `json:"winddir16Point"`
	WinddirDegree    string     `json:"winddirDegree"`
	WindspeedKmph    string     `json:"windspeedKmph"`
	WindspeedMiles   string     `json:"windspeedMiles"`
}

type NearestAreaDTO struct {
	AreaName   []ValueDTO `json:"areaName"`
	Country    []ValueDTO `json:"country"`
	Latitude   string     `json:"latitude"`
	Longitude  string     `json:"longitude"`
	Population string     `json:"population"`
	Region     []ValueDTO `json:"region"`
	WeatherURL []ValueDTO `json:"weatherUrl"`
}

type RequestDTO struct {
	Query string `json:"query"`
	Type  string `json:"type"`
}

type AstronomyDTO struct {
	MoonIllumination string `json:"moon_illumination"`
	MoonPhase        string `json:"moon_phase"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
}

type HourlyDTO struct {
	DewPointC        string     `json:"DewPointC"`
	DewPointF        string     `json:"DewPointF"`
	FeelsLikeC       string     `json:"FeelsLikeC"`
	FeelsLikeF       string     `json:"FeelsLikeF"`
	HeatIndexC       string     `json:"HeatIndexC"`
	HeatIndexF       string     `json:"HeatIndexF"`
	WindChillC       string     `json:"WindChillC"`
	WindChillF       string     `json:"WindChillF"`
	WindGustKmph     string     `json:"WindGustKmph"`
	WindGustMiles    string     `json:"WindGustMiles"`
	Chanceoffog      string     `json:"chanceoffog"`
	Chanceoffrost    string     `json:"chanceoffrost"`
	Chanceofhightemp string     `json:"chanceofhightemp"`
	Chanceofovercast string     `json:"chanceofovercast"`
	Chanceofrain     string     `json:"chanceofrain"`
	Chanceofremdry   string     `json:"chanceofremdry"`
	Chanceofsnow     string     `json:"chanceofsnow"`
	Chanceofsunshine string     `json:"chanceofsunshine"`
	Chanceofthunder  string     `json:"chanceofthunder"`
	Chanceofwindy    string     `json:"chanceofwindy"`
	Cloudcover       string     `json:"cloudcover"`
	DiffRad          string     `json:"diffRad"`
	Humidity         string     `json:"humidity"`
	LangPt           []ValueDTO `json:"lang_pt"`
	PrecipInches     string     `json:"precipInches"`
	PrecipMM         string     `json:"precipMM"`
	Pressure         string     `json:"pressure"`
	PressureInches   string     `json:"pressureInches"`
	ShortRad         string     `json:"shortRad"`
	TempC            string     `json:"tempC"`
	TempF            string     `json:"tempF"`
	Time             string     `json:"time"`
	UvIndex          string     `json:"uvIndex"`
	Visibility       string     `json:"visibility"`
	VisibilityMiles  string     `json:"visibilityMiles"`
	WeatherCode      string     `json:"weatherCode"`
	WeatherDesc      []ValueDTO `json:"weatherDesc"`
	WeatherIconURL   []ValueDTO `json:"weatherIconUrl"`
	Winddir16Point   string     `json:"winddir16Point"`
	WinddirDegree    string     `json:"winddirDegree"`
	WindspeedKmph    string     `json:"windspeedKmph"`
	WindspeedMiles   string     `json:"windspeedMiles"`
}

type WeatherDTO struct {
	Astronomy   []AstronomyDTO `json:"astronomy"`
	AvgtempC    string         `json:"avgtempC"`
	AvgtempF    string         `json:"avgtempF"`
	Date        string         `json:"date"`
	Hourly      []HourlyDTO    `json:"hourly"`
	MaxtempC    string         `json:"maxtempC"`
	MaxtempF    string         `json:"maxtempF"`
	MintempC    string         `json:"mintempC"`
	MintempF    string         `json:"mintempF"`
	SunHour     string         `json:"sunHour"`
	TotalSnowCm string         `json:"totalSnow_cm"`
	UvIndex     string         `json:"uvIndex"`
}

type WeatherResultDTO struct {
	CurrentCondition []CurrentConditionDTO `json:"current_condition"`
	NearestArea      []NearestAreaDTO      `json:"nearest_area"`
	Request          []RequestDTO          `json:"request"`
	Weather          []WeatherDTO          `json:"weather"`
}
