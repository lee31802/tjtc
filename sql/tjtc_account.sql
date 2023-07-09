CREATE TABLE tjtc_account.student_account (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    password varchar(64) Not Null,
    name     VARCHAR(8) NOT NULL,
    class_no TINYINT    NOT NULL,
    ct BIGINT NOT NULL
    );