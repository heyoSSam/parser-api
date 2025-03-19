package schema

var Skeleton = `
	CREATE TABLE IF NOT EXISTS codes (
		code_id SERIAL PRIMARY KEY,
		code VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATE,
		previous_date DATE
	);
	
	CREATE TABLE IF NOT EXISTS part (
		part_id SERIAL PRIMARY KEY,
		code_id INTEGER REFERENCES codes(code_id) ON DELETE CASCADE,
		code VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATE,
		previous_date DATE
	);
	
	CREATE TABLE IF NOT EXISTS sections (
		section_number SERIAL PRIMARY KEY,
		code_id INTEGER REFERENCES codes(code_id) ON DELETE CASCADE,
		part_id INTEGER REFERENCES part(part_id) ON DELETE CASCADE,
		section VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATE,
		previous_date DATE
	);
	
	CREATE TABLE IF NOT EXISTS paragraphs (
		paragraph_id SERIAL PRIMARY KEY,
		code_id INTEGER REFERENCES codes(code_id) ON DELETE CASCADE,
		part_id INTEGER REFERENCES part(part_id) ON DELETE CASCADE,
		section_number INTEGER REFERENCES sections(section_number) ON DELETE CASCADE,
		paragraph VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATE,
		previous_date DATE
	);
	
	CREATE TABLE IF NOT EXISTS articles (
		article_number SERIAL PRIMARY KEY,
		code_id INTEGER REFERENCES codes(code_id) ON DELETE CASCADE,
		part_id INTEGER REFERENCES part(part_id) ON DELETE CASCADE,
		section_number INTEGER REFERENCES sections(section_number) ON DELETE CASCADE,
		paragraph_id INTEGER REFERENCES paragraphs(paragraph_id) ON DELETE CASCADE,
		article VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATE,
		previous_date DATE
	);
	
	CREATE TABLE IF NOT EXISTS clauses (
		clause_number SERIAL PRIMARY KEY,
		code_id INTEGER REFERENCES codes(code_id) ON DELETE CASCADE,
		part_id INTEGER REFERENCES part(part_id) ON DELETE CASCADE,
		section_number INTEGER REFERENCES sections(section_number) ON DELETE CASCADE,
		paragraph_id INTEGER REFERENCES paragraphs(paragraph_id) ON DELETE CASCADE,
		article_number INTEGER REFERENCES articles(article_number) ON DELETE CASCADE,
		clause VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATE,
		previous_date DATE
	);
`
