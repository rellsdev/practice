SELECT 
    m.firstName AS manager_first_name, 
    m.lastName AS manager_last_name,
    s.firstName AS sales_first_name,
    SUM(od.quantityOrdered * od.priceEach) AS total_pendapatan
FROM employees s                                     
JOIN employees m ON s.reportsTo = m.employeeNumber   
JOIN customers c ON c.salesRepEmployeeNumber = s.employeeNumber
JOIN orders o ON o.customerNumber = c.customerNumber
JOIN orderdetails od ON od.orderNumber = o.orderNumber
GROUP BY s.employeeNumber, m.employeeNumber          
ORDER BY total_pendapatan ASC                        
LIMIT 1;

SELECT c1.customerName, c1.country, c1.creditLimit
FROM customers c1
WHERE c1.creditLimit > (
    SELECT AVG(c2.creditLimit)
    FROM customers c2
    WHERE c2.country = c1.country
);

WITH customer_totals AS (
    SELECT 
        c.customerNumber,
        c.customerName,
        c.country,
        SUM(p.amount) AS totalPayment
    FROM customers c
    JOIN payments p ON c.customerNumber = p.customerNumber
    GROUP BY c.customerNumber, c.customerName, c.country
)
SELECT 
    ct.customerName,
    ct.country,
    ct.totalPayment
FROM customer_totals ct
WHERE ct.totalPayment > (
    SELECT AVG(ct_sub.totalPayment)
    FROM customer_totals ct_sub
    WHERE ct_sub.country = ct.country
)
ORDER BY 
    ct.country ASC,
    ct.totalPayment DESC;

SELECT Customer_T.CustomerID, CustomerName, CustomerAddress, CustomerCity, CustomerState, CustomerPostalCode, 
       Order_T.OrderID, OrderDate, OrderedQuantity, ProductDescription, StandardPrice, 
       (OrderedQuantity * ProductStandardPrice) 
FROM Customer_T, Order_T, OrderLine_T, Product_T
WHERE Order_T.CustomerID = Customer_T.CustomerID
  AND Order_T.OrderID = OrderLine_T.OrderID
  AND OrderLine_T.ProductID = Product_T.ProductID
  AND Order_T.OrderID = 1006; 

/*Evaluasi Produk dengan Harga di Atas Rata-rata
Tim inventory ingin mengidentifikasi produk-produk dengan harga beli yang 
lebih tinggi dibandingkan rata-rata harga beli seluruh produk.
Tampilkan:
• productCode
• productName
• productLine
• buyPrice
Hanya tampilkan produk yang memiliki buyPrice lebih besar dari rata-rata 
buyPrice seluruh produk.
Urutkan berdasarkan buyPrice dari yang tertinggi.
Gunakan tabel:
• products*/

SELECT productCode, productName, productLine, buyPrice
FROM products 
WHERE buyPrice > (SELECT AVG(buyPrice) FROM products)
ORDER by buyPrice DESC;


/*Analisis Aktivitas Pelanggan dan Pembayaran
Departemen keuangan ingin mengetahui kontribusi setiap pelanggan terhadap perusahaan.
Tampilkan:
• customerName
• Jumlah pesanan yang pernah dibuat (totalOrders)
• Total pembayaran yang pernah dilakukan (totalPayments)
Urutkan berdasarkan totalPayments dari yang terbesar.
Gunakan tabel:
• customers
• orders
• payments*/

SELECT c.customerName, COUNT(DISTINCT o.orderNumber) AS totalOrders, 
SUM(p.amount) AS totalPayments
FROM customers c
JOIN orders o
ON c.customerNumber = o.customerNumber
JOIN payments p
ON c.customerNumber = p.customerNumber
GROUP BY c.customerName
ORDER BY totalPayments DESC;



/*Identifikasi Kantor dengan Jumlah Karyawan Tinggi
Manajemen ingin mengetahui kantor cabang yang memiliki jumlah karyawan lebih banyak
dibandingkan rata-rata jumlah karyawan per kantor.
Tampilkan:
• officeCode
• city
• country
• Jumlah karyawan pada kantor tersebut (totalEmployees)
Hanya tampilkan kantor yang memiliki jumlah karyawan di atas rata-rata jumlah karyawan seluruh
kantor. Urutkan berdasarkan totalEmployees secara menurun.
Gunakan tabel:
• offices
• employees*/

SELECT ofc.officeCode, ofc.city, ofc.country, COUNT(e.employeeNumber) AS totalEmployees
FROM offices ofc
JOIN employees e
ON ofc.officeCode = e.officeCode
GROUP BY ofc.officeCode, ofc.city, ofc.country
HAVING COUNT(e.employeeNumber) > (
	SELECT AVG(total_employee) #lapisan Kedua
	FROM (SELECT officeCode, COUNT(*) AS total_employee #lapisan Ketiga
		FROM employees
		GROUP BY officeCode)AS rataRata
	)
ORDER BY totalEmployees DESC;


WHERE -> GROUP BY -> HAVING 


/*Analisis Produk Paling Diminati
Tim pemasaran ingin mengetahui produk yang telah terjual dalam jumlah besar.
Tampilkan:
• productCode
• productName
• Total unit terjual (totalQuantitySold)
• Total pendapatan produk (totalRevenue)
Hanya tampilkan produk yang memiliki total unit terjual lebih besar dari 500 unit.
Urutkan berdasarkan totalRevenue dari yang terbesar.
Gunakan tabel:
• products
• orderdetails*/

SELECT pr.productCode, pr.productName, SUM(od.quantityOrdered) AS totalQuantitySold,
SUM(od.quantityOrdered * od.priceEach) AS totalRevenue
FROM products pr
JOIN orderdetails od
ON pr.productCode = od.productCode
GROUP BY pr.productCode, pr.productName
HAVING SUM(od.quantityOrdered) > 500
ORDER BY totalRevenue DESC;