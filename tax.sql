CREATE TABLE IF NOT EXISTS fed_tax_table (
	pay_freq CHAR(1),
    marital CHAR(2),
    earning Decimal(10,2),
    amount Decimal(10,2),
    percentage Decimal(6,2),
	adjust_income Decimal(10,2)

);
	Delete from fed_tax_table where 1=1; 


	INSERT INTO fed_tax_table VALUES ('A', 'M', 16300, 0, 0, 12900);
	INSERT INTO fed_tax_table VALUES ('A', 'M', 39500, 0, 0.1, 12900);
	INSERT INTO fed_tax_table VALUES ('A', 'M', 110600, 2320.00, 0.12, 12900);
	INSERT INTO fed_tax_table VALUES ('A', 'M', 217350, 10852.00, 0.22, 12900);
	INSERT INTO fed_tax_table VALUES ('A', 'M', 400200, 34337.00, 0.24, 12900);
	INSERT INTO fed_tax_table VALUES ('A', 'M', 503750, 78221.00, 0.32, 12900);
	INSERT INTO fed_tax_table VALUES ('A', 'M', 747500, 111357.00, 0.35, 12900);
	INSERT INTO fed_tax_table VALUES ('A', 'M', 9999999, 196669.50, 0.37, 12900);

	--Input federal tax rate for S
	INSERT INTO fed_tax_table VALUES ('A', 'S', 6000, 0, 0, 8600);
	INSERT INTO fed_tax_table VALUES ('A', 'S', 17600, 0, 0.1, 8600);
	INSERT INTO fed_tax_table VALUES ('A', 'S', 53150, 1160, 0.12, 8600);
	INSERT INTO fed_tax_table VALUES ('A', 'S', 106525, 5426, 0.22, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'S', 197950, 17168.50, 0.24, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'S', 249725, 39110.50, 0.32, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'S', 615350, 55678.50, 0.35, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'S', 9999999, 183647.25, 0.37, 8600);

	--Input federal tax rate for HH
	INSERT INTO fed_tax_table VALUES ( 'A', 'HH', 13300, 0, 0, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'HH', 29850, 0, 0.1, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'HH', 76400, 1655.00, 0.12, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'HH', 113800, 7241.00, 0.22, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'HH', 205250, 15469.00, 0.24, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'HH', 257000, 37417.00, 0.32, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'HH', 622650, 53977.00, 0.35, 8600);
	INSERT INTO fed_tax_table VALUES ( 'A', 'HH', 9999999, 181954.50, 0.37, 8600);

	--Input federal tax rate for 2M
	INSERT INTO fed_tax_table VALUES ( 'A', '2M', 14600, 0, 0, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2M', 26200, 0, 0.1, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2M', 61750, 1160, 0.12, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2M', 115125, 5426, 0.22, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2M', 206550, 17168.50, 0.24, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2M', 258325, 39110.50, 0.32, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2M', 380200, 55678.50, 0.35, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2M', 9999999, 98334.75, 0.37, 0);

	--Input federal tax rate for 2S
	INSERT INTO fed_tax_table VALUES ( 'A', '2S', 7300, 0, 0, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2S', 13100, 0, 0.1, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2S', 30875, 580, 0.12, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2S', 57563, 2713, 0.22, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2S', 103275, 8584.25, 0.24, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2S', 129163, 19555.25, 0.32, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2S', 311975, 27839.25, 0.35, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2S', 9999999, 91823.63, 0.37, 0);

	--Input federal tax rate for 2H
	INSERT INTO fed_tax_table VALUES ( 'A', '2H', 10950, 0, 0, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2H', 19225, 0, 0.1, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2H', 42500, 827.50, 0.12, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2H', 61200, 3620.50, 0.22, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2H', 106925, 7734.50, 0.24, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2H', 132800, 18708.50, 0.32, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2H', 315625, 26988.50, 0.35, 0);
	INSERT INTO fed_tax_table VALUES ( 'A', '2H', 9999999, 90977.25, 0.37, 0);
