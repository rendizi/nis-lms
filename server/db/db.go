package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "lms"
)

var Db *sql.DB

func Init() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to PostgreSQL!")

	//Table Users
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS students (
			id SERIAL PRIMARY KEY,
			login TEXT NOT NULL UNIQUE, 
			password TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			klass TEXT NOT NULL,
			parallel INTEGER NOT NULL,
			school TEXT NOT NULL,
			
			solved INTEGER NOT NULL,
			leetcode TEXT,
			badges TEXT,
			rating INTEGER NOT NULL
);
    `
	_, err = Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'students' created successfully!")

	//Table Expressions
	createTableQuery = `
        CREATE TABLE IF NOT EXISTS teachers (
    		id SERIAL PRIMARY KEY,
    		login TEXT NOT NULL UNIQUE,
    		password TEXT NOT NULL,
    		email TEXT NOT NULL,
    		school TEXT NOT NULL,
    		leetcode TEXT
                                            );

    `
	_, err = Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'students' created successfully!")

	//Table Operations
	createTableQuery = `
        CREATE TABLE IF NOT EXISTS classwork (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			teacher_id INTEGER NOT NULL,
			tasks_id INTEGER[],
			deadline TEXT, 
			FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
);
    `
	_, err = Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'classwork' created successfully!")

	createTableQuery = `
		CREATE TABLE IF NOT EXISTS classwork_students (
    classwork_id INTEGER NOT NULL,
    student_id INTEGER NOT NULL,
    PRIMARY KEY (classwork_id, student_id),
    FOREIGN KEY (classwork_id) REFERENCES classwork(id) ON DELETE CASCADE,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
);


`
	_, err = Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'classwork_students' created successfully!")

	createTableQuery = `
		CREATE TABLE IF NOT EXISTS tasks (
    		id SERIAL PRIMARY KEY,
    		title TEXT NOT NULL,
    		description TEXT NOT NULL,
    		author TEXT NOT NULL,
    		difficulty TEXT NOT NULL,
    		tests TEXT NOT NULL,
    		image TEXT,
    		firstExample TEXT NOT NULL,
    		secondExample TEXT,
    		thirdExample TEXT
		                                 );
`
	_, err = Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'tasks' created successfully!")

	createTableQuery = `
		CREATE TABLE IF NOT EXISTS solutions (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    task_id INT NOT NULL,
    solution TEXT NOT NULL,
    submission_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_student
        FOREIGN KEY(student_id) 
        REFERENCES students(id),
    CONSTRAINT fk_task
        FOREIGN KEY(task_id) 
        REFERENCES tasks(id),
    UNIQUE(student_id, task_id) 
);

`
	_, err = Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'solutions' created successfully!")
}
