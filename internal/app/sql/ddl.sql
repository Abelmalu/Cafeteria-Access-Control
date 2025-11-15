-- sql/ddl.sql


-- Table for different cafeterias
CREATE TABLE IF NOT EXISTS cafeterias (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,  -- Unique ID for the location
careteria_name VARCHAR(100) NOT NULL,
careteria_location VARCHAR(100) NOT NULL
);



-- Table for batches in university 
CREATE TABLE IF NOT EXISTS batches(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,  -- Cleaned up line break
    batches_name VARCHAR(100) NOT NULL,
    cafeteria_id INT NOT NULL,
    FOREIGN KEY (cafeteria_id) REFERENCES cafeterias(id)
);

-- Table for different devices used for scanning 
CREATE TABLE IF NOT EXISTS devices(
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
`name`  VARCHAR(100) NOT NULL,
serial_number VARCHAR(100) NOT NULL,
cafeteria_id INT NOT NULL,
FOREIGN KEY (cafeteria_id) REFERENCES cafeterias(id)
);



-- Table of registered meals and the time window 
CREATE TABLE IF NOT EXISTS meals(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    `name` VARCHAR(100) NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL
);



-- Table for registered students
CREATE TABLE IF NOT EXISTS students (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,  
first_name VARCHAR(100) NOT NULL,
middle_name VARCHAR(100) NOT NULL,
last_name VARCHAR(100) NOT NULL,
rfid_tag VARCHAR(100) NOT NULL,
image_url VARCHAR(100) NOT NULL,
batch_id INT NOT NULL,
FOREIGN KEY (batch_id) REFERENCES batches (id)

);




-- Table to log all access attempts (successful or failed)
CREATE TABLE IF NOT EXISTS meal_access_log (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    
    scan_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) NOT NULL,
    student_id INT NOT NULL,
    cafeteria_id INT NOT NULL,
    meal_id INT NOT NULL,
    device_id INT NOT NULL,

    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (cafeteria_id) REFERENCES cafeterias(id),
    FOREIGN KEY (meal_id) REFERENCES meals(id),
    FOREIGN KEY (device_id) REFERENCES devices(id)
);
