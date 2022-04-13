from flask import Flask
from flask import jsonify
app = Flask(__name__)
import grpc

import sys
sys.path.append("./countries")
sys.path.append("./helloworld")

from countries import countries_pb2
from countries import countries_pb2_grpc

from helloworld import helloworld_pb2
from helloworld import helloworld_pb2_grpc

def obj_to_dict(obj):return obj.__dict__

@app.route('/countries/<countrie>')
def countrie(countrie): 
    try:
        channel = grpc.insecure_channel("grpc-server-go:8000")
        stub = countries_pb2_grpc.CountryStub(channel)
        countryRequest = countries_pb2.CountryRequest(name=countrie)
        countryResponse = stub.Search(countryRequest)
        return jsonify({
            "name": countryResponse.name,
            "alpha2Code": countryResponse.alpha2Code,
            "capital": countryResponse.capital,
            "subregion": countryResponse.subregion,
            "population": countryResponse.population,
            "nativeName": countryResponse.nativeName
        })
    except Exception as e:
      print(e)
      return e


@app.route('/helloworld/<name>')
def helloWorld(name): 
    try:
        channel = grpc.insecure_channel("grpc-server-go:8000")
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        helloRequest = helloworld_pb2.HelloRequest(name=name)
        helloResponse = stub.SayHello(helloRequest)
        return jsonify({
            "message": helloResponse.message
        })
    except Exception as e:
      print(e)
      return e

if __name__ == '__main__':
    app.run(host = "0.0.0.0", port=6000)