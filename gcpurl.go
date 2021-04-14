package gcpurl

import (
	"fmt"
	"regexp"
)

type Region string

// https://cloud.google.com/compute/docs/regions-zones
const (
	AsiaEast1Region              Region = "asia-east1"
	AsiaEast2Region              Region = "asia-east2"
	AsiaNorthEast1Region         Region = "asia-northeast1"
	AsiaNorthEast2Region         Region = "asia-northeast2"
	AsiaNorthEast3Region         Region = "asia-northeast3"
	AsiaSouth1Region             Region = "asia-south1"
	AsiaSouthEast1Region         Region = "asia-southeast1"
	AsiaSouthEast2Region         Region = "asia-southeast2"
	AustraliaSouthEast1Region    Region = "australia-southeast1"
	EuropeCentral2Region         Region = "europe-central2"
	EuropeNorth1Region           Region = "europe-north1"
	EuropeWest1Region            Region = "europe-west1"
	EuropeWest2Region            Region = "europe-west2"
	EuropeWest3Region            Region = "europe-west3"
	EuropeWest4Region            Region = "europe-west4"
	EuropeWest6Region            Region = "europe-west6"
	NorthAmericaNorthEast1Region Region = "northamerica-northeast1"
	SouthAmericaEast1Region      Region = "southamerica-east1"
	UsCentral1Region             Region = "us-central1"
	UsEast1Region                Region = "us-east1"
	UsEast4Region                Region = "us-east4"
	UsWest1Region                Region = "us-west1"
	UsWest2Region                Region = "us-west2"
	UsWest3Region                Region = "us-west3"
	UsWest4Region                Region = "us-west4"
)

func (r Region) Zones() []string {
	switch r {
	case AsiaEast1Region:
		return []string{"a", "b", "c"}
	case AsiaEast2Region:
		return []string{"a", "b", "c"}
	case AsiaNorthEast1Region:
		return []string{"a", "b", "c"}
	case AsiaNorthEast2Region:
		return []string{"a", "b", "c"}
	case AsiaNorthEast3Region:
		return []string{"a", "b", "c"}
	case AsiaSouth1Region:
		return []string{"a", "b", "c"}
	case AsiaSouthEast1Region:
		return []string{"a", "b", "c"}
	case AsiaSouthEast2Region:
		return []string{"a", "b", "c"}
	case AustraliaSouthEast1Region:
		return []string{"a", "b", "c"}
	case EuropeCentral2Region:
		return []string{"a", "b", "c"}
	case EuropeNorth1Region:
		return []string{"a", "b", "c"}
	case EuropeWest1Region:
		return []string{"b", "c", "d"}
	case EuropeWest2Region:
		return []string{"a", "b", "c"}
	case EuropeWest3Region:
		return []string{"a", "b", "c"}
	case EuropeWest4Region:
		return []string{"a", "b", "c"}
	case EuropeWest6Region:
		return []string{"a", "b", "c"}
	case NorthAmericaNorthEast1Region:
		return []string{"a", "b", "c"}
	case SouthAmericaEast1Region:
		return []string{"a", "b", "c"}
	case UsCentral1Region:
		return []string{"a", "b", "c", "f"}
	case UsEast1Region:
		return []string{"b", "c", "d"}
	case UsEast4Region:
		return []string{"a", "b", "c"}
	case UsWest1Region:
		return []string{"a", "b", "c"}
	case UsWest2Region:
		return []string{"a", "b", "c"}
	case UsWest3Region:
		return []string{"a", "b", "c"}
	case UsWest4Region:
		return []string{"a", "b", "c"}
	}
	return []string{}
}

var (
	Regions = []Region{
		AsiaEast1Region,
		AsiaEast2Region,
		AsiaNorthEast1Region,
		AsiaNorthEast2Region,
		AsiaNorthEast3Region,
		AsiaSouth1Region,
		AsiaSouthEast1Region,
		AsiaSouthEast2Region,
		AustraliaSouthEast1Region,
		EuropeCentral2Region,
		EuropeNorth1Region,
		EuropeWest1Region,
		EuropeWest2Region,
		EuropeWest3Region,
		EuropeWest4Region,
		EuropeWest6Region,
		NorthAmericaNorthEast1Region,
		SouthAmericaEast1Region,
		UsCentral1Region,
		UsEast1Region,
		UsEast4Region,
		UsWest1Region,
		UsWest2Region,
		UsWest3Region,
		UsWest4Region,
	}
	parsePatterns = map[Region]*regexp.Regexp{}
)

var (
	ErrParseURL = fmt.Errorf("gcpurl: failed to parse url")
)

func init() {
	for _, region := range Regions {
		pat := fmt.Sprintf(`%s-([a-z][a-z0-9-]+)\.(.+)$`, string(region))
		parsePatterns[region] = regexp.MustCompile(pat)
	}
}

type URL struct {
	Host          string
	Region        Region
	ProjectID     string
	ServiceDomain string
}

func ParseURL(url string) (*URL, error) {
	for _, region := range Regions {
		matches := parsePatterns[region].FindAllStringSubmatch(url, -1)
		if len(matches) == 0 {
			continue
		}
		if len(matches[0]) != 3 {
			continue
		}
		return &URL{
			Host:          matches[0][0],
			Region:        region,
			ProjectID:     matches[0][1],
			ServiceDomain: matches[0][2],
		}, nil
	}
	return nil, ErrParseURL
}
