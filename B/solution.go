package main

import (
	"fmt"
)

type Result struct {
	GoodsId   int
	GoodsName string
}

const solutionQuery = `
	WITH all_tags AS (
	    SELECT 
		   COUNT(1) AS count 
	    FROM
			tags
	)
	SELECT
		goods.id,
		goods.name
	FROM (
		SELECT
			goods_id,
			COUNT(tags_goods.tag_id) AS tags_count
		FROM
			tags_goods,
		    all_tags
	    GROUP BY 
			goods_id
		HAVING
			tags_count = all_tags.count
	) AS goods_with_all_tags
	INNER JOIN goods ON goods.id = goods_with_all_tags.goods_id
`

func Solution() ([]Result, error) {
	db, err := NewDb()
	if err != nil {
		return nil, fmt.Errorf("NewDb(): %w", err)
	}
	rows, err := db.Query(solutionQuery)
	if err != nil {
		return nil, fmt.Errorf("db.Query(\"%v\"): %w", solutionQuery, err)
	}
	defer rows.Close()

	results := make([]Result, 0, 10)
	var i int

	for rows.Next() {
		results = append(results, Result{})
		err := rows.Scan(&results[i].GoodsId, &results[i].GoodsName)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan(): %w", err)
		}
		i++
	}

	return results, nil
}
