#include <chrono>
#include <iostream>
#include <string>
#include <thread>
#include <zmq.hpp>
#include "EmployeeTest.pb.h"
#include "TopData.pb.h"
using namespace EmployeeTest;
using std::cout;
using std::endl;
using std::ends;

int main() {
  using namespace std::chrono_literals;

  // initialize the zmq context with a single IO thread
  zmq::context_t context{1};

  // construct a REP (reply) socket and bind to interface
  zmq::socket_t socket{context, zmq::socket_type::rep};
  socket.bind("tcp://*:9527");

  // prepare some static data for responses
  const std::string data{"World"};

  for (;;) {
    zmq::message_t request;

    // receive a request from client
    socket.recv(request, zmq::recv_flags::none);
    //std::cout << "Received " << request.to_string() << std::endl;

    TrZeroMQMsg::TopData tempTopData;
    //反序列化
    tempTopData.ParseFromString(request.to_string());
    std::cout << tempTopData.type() << std::endl;

    //解析的数据
    EmployeeTest::Employee getMyData;
    //反序列化
    getMyData.ParseFromString(tempTopData.rawdata());
    cout << getMyData.id() << ends << getMyData.age() << ends
         << getMyData.name() << endl;

    // simulate work
    std::this_thread::sleep_for(1s);

    // send the reply to the client
    socket.send(zmq::buffer(data), zmq::send_flags::none);
  }

  return 0;
}
