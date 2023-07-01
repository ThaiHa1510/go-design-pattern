# Observer pattern
Observer pattern là một trong những pattern thuộc nhóm Behavioral Pattern. Nó cho phép các đối tượng theo dõi và tự động cập nhật thông tin khi trạng thái của một đối tượng khác thay đổi. Pattern này được sử dụng để giảm sự phụ thuộc giữa các đối tượng và giúp tăng tính tái sử dụng và mở rộng của code..

# Một số khái niệm cơ bản:
- Subject: Đối tượng chủ đề là đối tượng mà cac Observer quan sát , khi trạng thái của Subject thay đổi , tất cả các đối tượng Observer khác sẽ được thông báo
- Observer : Đối tượng quan sát . Là đối tượng sẽ được thông báo những thay đổi của đối tượng Subject. Đối tượng Observer phải được thêm vào đối tượng Subject và được có một phương thức chung để được thông báo 

# Một số trường hợp sử dụng cụ thể 
- Hệ thống đăng ký sự kiện: Observer Pattern được sử dụng rất phổ biến trong các hệ thống đăng ký sự kiện. Khi người dùng đăng ký một sự kiện, hệ thống sẽ tạo ra một đối tượng Observer, sau đó sẽ gửi thông tin sự kiện đến đối tượng này khi sự kiện xảy ra. Đối tượng Observer sẽ tiếp nhận thông tin và thực hiện các hành động tương ứng với sự kiện đó.

- Hệ thống thông báo: Observer Pattern cũng được sử dụng để thiết kế các hệ thống thông báo. Ví dụ như trong một ứng dụng thương mại điện tử, khi khách hàng đặt hàng thành công, hệ thống sẽ thông báo cho khách hàng về đơn hàng của họ bằng cách gửi email hoặc thông báo trên ứng dụng di động. Hệ thống thông báo sẽ được thiết kế dựa trên Observer Pattern, với đối tượng Observer là khách hàng và đối tượng Subject là đơn hàng.

- Hệ thống theo dõi trạng thái: Trong các hệ thống theo dõi trạng thái, Observer Pattern cũng được sử dụng để thông báo cho người dùng về các sự kiện quan trọng. Ví dụ như trong một hệ thống theo dõi máy chủ, khi máy chủ gặp sự cố, hệ thống sẽ thông báo cho người quản trị thông qua email hoặc tin nhắn. Hệ thống này sẽ được thiết kế dựa trên Observer Pattern, với đối tượng Observer là người quản trị và đối tượng Subject là máy chủ.

- Hệ thống đa luồng: Observer Pattern cũng được sử dụng trong các ứng dụng đa luồng, khi các luồng cần tương tác với nhau để truyền tải thông tin hoặc thực hiện các hành động. Trong trường hợp này, các luồng sẽ được thiết kế dựa trên Observer Pattern, với đối tượng Observer là các luồng và đối tượng Subject là tác vụ cần thực hiện.