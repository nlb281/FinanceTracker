CREATE TABLE IF NOT EXISTS categories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE,
  type TEXT NOT NULL CHECK (type IN ('income', 'expense'))
);

CREATE TABLE IF NOT EXISTS transactions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  amount REAL NOT NULL CHECK (amount >= 0),
  type TEXT NOT NULL CHECK (type IN ('income', 'expense')),
  category_id INTEGER NOT NULL,
  date TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_transactions_date ON transactions(date);
CREATE INDEX IF NOT EXISTS idx_transactions_type ON transactions(type);
CREATE INDEX IF NOT EXISTS idx_transactions_category_id ON transactions(category_id);

INSERT OR IGNORE INTO categories (name, type) VALUES 
  -- Доходы
  ('Зарплата', 'income'),
  ('Подработка', 'income'),
  ('Бизнес', 'income'),
  ('Кэшбэк', 'income'),
  ('Подарки', 'income'),
  ('Возврат долгов', 'income'),
  ('Проценты по вкладам', 'income'),
  ('Продажа вещей', 'income'),
  ('Выплаты/Пособия', 'income'),
  
  -- Расходы
  ('Аренда', 'expense'),
  ('Коммунальные услуги', 'expense'),
  ('Связь и Интернет', 'expense'),
  ('Подписки', 'expense'),
  ('Кредиты', 'expense'),
  ('Продукты', 'expense'),
  ('Кафе и Рестораны', 'expense'),
  ('Транспорт', 'expense'),
  ('Здоровье', 'expense'),
  ('Одежда', 'expense'),
  ('Красота', 'expense'),
  ('Развлечения', 'expense'),
  ('Образование', 'expense'),
  ('Подарки', 'expense'),
  ('Дом и ремонт', 'expense'),
  ('Питомцы', 'expense'),
  ('Инвестиции', 'expense'),
  ('Накопления', 'expense'),
  ('Разное', 'expense');