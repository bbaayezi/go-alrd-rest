package dbquery

var (
	GetLatestPublications = `select title, unnest(author) as author, date
		from t_abstract_data
		order by date desc
		limit 10;
	`
	GetOverview = `select publications, authors, total_citations, start_year, end_year
		from (
			select count(title) as publications, count(distinct author) as authors,
				sum(citedby_count) as total_citations,
				split_part(min(date), '-', 1) as start_year, split_part(max(date), '-', 1) as end_year
			from t_abstract_data
		) t
		join lateral (
			select count(distinct author) as dauthor
			from (
				select unnest(author) as author
				from t_abstract_data
			) t
		) as se on true;
	`
	GetYearlySummary = `select left(date, 4) as year, count(title)
		from t_abstract_data
		group by year
		order by year desc
		limit 6;
	`
	GetPartnerCountry = `select distinct country as country, count(country) as number
		from (
			select unnest(country) as country
			from t_abstract_data
		) t
		group by country;
	`
	GetKeywords = `select distinct keywords as keywords, count(keywords) as number
		from (
			select unnest(subject_area) as keywords
			from t_abstract_data
		) t
		group by keywords
		order by number desc
		limit 50;
	`
	GetPublisher = `select publisher, count(publisher) as number
		from (
			select publisher as publisher
				from t_abstract_data
		) t
		where publisher != ''
		group by publisher
		order by number desc
		limit 10;
	`
	GetContentType = `select content_type, count(content_type) as number
		from t_abstract_data
		group by content_type
		order by number desc;
	`
)
