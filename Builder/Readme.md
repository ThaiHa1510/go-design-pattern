# Builder pattern
Builder pattern là một trong những pattern thuộc nhóm Creational Pattern. Nó cung cấp một cách tiếp cận khác để xây dựng đối tượng, tách rời việc xây dựng đối tượng khỏi việc biểu diễn của nó. Builder Pattern cho phép ta tạo ra các đối tượng phức tạp bằng cách sử dụng các bước riêng lẻ và độc lập với nhau.

# Các thành phần cơ bản:
- Builder: Định nghĩa các bước cần thiết để xây dựng đối tượng.
- ConcreteBuilder: Cài đặt các bước cụ thể để xây dựng đối tượng.
- Product: Đối tượng được xây dựng bởi ConcreteBuilder.
- Director: Quản lý việc thực thi các bước để tạo đối tượng, tạo ra một đối tượng sử dụng Builder.

# Một trường hợp sử dụng Builder Pattern
- Quá trình tạo đối tượng phức tạp, có nhiều bước và các bước này có thể thay đổi trong quá trình phát triển.
- Cần tái sử dụng các bước tạo đối tượng cho các đối tượng khác nhau.
- Cần tách rời quá trình tạo đối tượng và đối tượng được tạo ra để có thể dễ dàng thay đổi hoặc mở rộng.