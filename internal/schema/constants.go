package schema

var Skeleton = `
	CREATE TABLE IF NOT EXISTS codes (
		code_id INTEGER UNSIGNED PRIMARY KEY,
		code VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATETIME,
		previous_date DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS part (
		part_id INTEGER UNSIGNED PRIMARY KEY,
		code_id INTEGER UNSIGNED REFERENCES codes(code_id) ON DELETE CASCADE,
		code VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATETIME,
		previous_date DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS sections (
		section_number INTEGER UNSIGNED PRIMARY KEY,
		code_id INTEGER UNSIGNED REFERENCES codes(code_id) ON DELETE CASCADE,
		part_id INTEGER UNSIGNED REFERENCES part(part_id) ON DELETE CASCADE,
		section VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATETIME,
		previous_date DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS paragraphs (
		paragraph_id INTEGER UNSIGNED PRIMARY KEY,
		code_id INTEGER UNSIGNED REFERENCES codes(code_id) ON DELETE CASCADE,
		part_id INTEGER UNSIGNED REFERENCES part(part_id) ON DELETE CASCADE,
		section_number INTEGER UNSIGNED REFERENCES sections(section_number) ON DELETE CASCADE,
		paragraph VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATETIME,
		previous_date DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS articles (
		article_number INTEGER UNSIGNED PRIMARY KEY,
		code_id INTEGER UNSIGNED REFERENCES codes(code_id) ON DELETE CASCADE,
		part_id INTEGER UNSIGNED REFERENCES part(part_id) ON DELETE CASCADE,
		section_number INTEGER UNSIGNED REFERENCES sections(section_number) ON DELETE CASCADE,
		paragraph_id INTEGER UNSIGNED REFERENCES paragraphs(paragraph_id) ON DELETE CASCADE,
		article VARCHAR(255),
		version INTEGER,
		status BOOLEAN,
		actual_date DATETIME,
		previous_date DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS clauses (
		clause_id INTEGER UNSIGNED PRIMARY KEY,
		clause_number INTEGER UNSIGNED,
		code_id INTEGER UNSIGNED REFERENCES codes(code_id) ON DELETE CASCADE,
		part_id INTEGER UNSIGNED REFERENCES part(part_id) ON DELETE CASCADE,
		section_number INTEGER UNSIGNED REFERENCES sections(section_number) ON DELETE CASCADE,
		paragraph_id INTEGER UNSIGNED REFERENCES paragraphs(paragraph_id) ON DELETE CASCADE,
		article_number INTEGER UNSIGNED REFERENCES articles(article_number) ON DELETE CASCADE,
		clause TEXT,
		version INTEGER,
		status BOOLEAN,
		actual_date DATETIME,
		previous_date DATETIME
	);

	INSERT INTO codes (code_id, code, version, status, actual_date, previous_date)
	VALUES (1, 'No Reference', 1, TRUE, NOW(), NULL);

	INSERT INTO part (part_id, code_id, code, version, status, actual_date, previous_date) 
	VALUES (0, 1, 'No Reference', 1, TRUE, NOW(), NULL);

	INSERT INTO sections (section_number, code_id, part_id, section, version, status, actual_date, previous_date) 
	VALUES (0, 1, 0, 'No Reference', 1, TRUE, NOW(), NULL);

	INSERT INTO paragraphs (paragraph_id, code_id, part_id, section_number, paragraph, version, status, actual_date, previous_date) 
	VALUES (0, 1, 0, 0, 'No Reference', 1, TRUE, NOW(), NULL);

	INSERT INTO articles (article_number, code_id, part_id, section_number, paragraph_id, article, version, status, actual_date, previous_date) 
	VALUES (0, 1, 0, 0, 0, 'No Reference', 1, TRUE, NOW(), NULL);

	INSERT INTO clauses (clause_id, clause_number, code_id, part_id, section_number, paragraph_id, article_number, clause, version, status, actual_date, previous_date) 
	VALUES (0, 0, 1, 0, 0, 0, 0, 'No Reference', 1, TRUE, NOW(), NULL);
`
