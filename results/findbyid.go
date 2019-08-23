package results

import (
	"encoding/json"
)

/*
"101748479": {
    "id": 101748479,
    "name": "München",
    "placetype": "locality",
    "rank": {
      "min": 9,
      "max": 10
    },
    "population": 1260391,
    "lineage": [
      {
        "continent_id": 102191581,
        "country_id": 85633111,
        "county_id": 102063261,
        "locality_id": 101748479,
        "macrocounty_id": 404227567,
        "region_id": 85682571
      }
    ],
    "geom": {
      "area": 0.031614,
      "bbox": "11.382263,48.0634908248,11.7231562646,48.2282863287",
      "lat": 48.152126,
      "lon": 11.544467
    },
    "names": {
      "afr": [
        "München"
      ],
*/

type FindByIDResults map[string]interface{}

func (s *FindByIDResults) Results() []interface{} {

     results := make([]interface{}, 0)
     
     for _, v := range *s {
     	 results = append(results, v)
     }
	 
     return results
}

func NewFindByIDResults(body []byte) (*FindByIDResults, error) {

	var id_results *FindByIDResults
	err := json.Unmarshal(body, &id_results)

	if err != nil {
		return nil, err
	}

	return id_results, nil
}
