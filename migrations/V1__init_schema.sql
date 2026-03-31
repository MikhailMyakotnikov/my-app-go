CREATE TABLE `students` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `teachers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `courses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `teacher_id` int NOT NULL,
  `duration` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `Course_Teacher_FK` (`teacher_id`),
  CONSTRAINT `Course_Teacher_FK` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE `students_courses` (
  `student_id` int NOT NULL,
  `course_id` int NOT NULL,
  KEY `StudentCourse_Student_FK` (`student_id`),
  KEY `StudentCourse_Course_FK` (`course_id`),
  CONSTRAINT `StudentCourse_Course_FK` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `StudentCourse_Student_FK` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);