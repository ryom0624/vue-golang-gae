DROP TABLE IF EXISTS articles;

CREATE TABLE articles (
  id integer AUTO_INCREMENT,
  name varchar(255),
  description text,
  primary key(id)
);

INSERT INTO articles (id, name, description)
VALUES
  (1, 'サイトリニューアルしました', 'ほげほげほげほげ');

INSERT INTO articles (id, name, description)
VALUES
  (2, 'CIを刷新いたしました', 'ふがふがふがふが');

INSERT INTO articles (id, name, description)
VALUES
  (3, '東京〇〇にブース出展いたしました。', 'ぴよぴよぴよぴよ');