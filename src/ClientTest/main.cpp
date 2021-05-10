#include <iostream>
#include <string>

#include "EmployeeTest.pb.h"
#include "zmq.hpp"
#include "zmq_addon.hpp"
using std::cout;
using std::endl;
using std::ends;
using namespace EmployeeTest;

int main() {
  //发送的数据
  EmployeeTest::Employee myData;
  myData.set_id(918);
  myData.set_age(30);
  myData.set_name("test wang yi yun");
  //序列化
  auto tempData{myData.SerializeAsString()};
  cout << "the serialize data size:" << tempData.size() << endl;

  //解析的数据
  EmployeeTest::Employee getMyData;
  //反序列化
  getMyData.ParseFromString(tempData);
  cout << getMyData.id() << ends << getMyData.age() << ends << getMyData.name()
       << endl;

  // initialize the zmq context with a single IO thread
  zmq::context_t context{1};

  // construct a REQ (request) socket and connect to interface
  zmq::socket_t socket{context, zmq::socket_type::req};
  socket.connect("tcp://localhost:5555");
  // socket.set(zmq::sockopt::subscribe, "");

  const std::string data{"Hello"};
  for (auto request_num = 0; request_num < 10; ++request_num) {
    // send the request message
    std::cout << "Sending Hello " << request_num << "..." << std::endl;
    //socket.send(zmq::buffer(data), zmq::send_flags::none);
    socket.send(zmq::buffer(tempData), zmq::send_flags::none);

    // wait for reply from server
    zmq::message_t reply{};
    socket.recv(reply, zmq::recv_flags::none);

    std::cout << "Received " << reply.to_string();
    std::cout << " (" << request_num << ")";
    std::cout << std::endl;
  }

  return 0;
}
