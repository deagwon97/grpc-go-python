# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: countries.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0f\x63ountries.proto\x12\tcountries\"\x1e\n\x0e\x43ountryRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"8\n\nCurrencies\x12\x0c\n\x04\x63ode\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06symbol\x18\x03 \x01(\t\"\xaa\x01\n\x0f\x43ountryResponse\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x12\n\nalpha2Code\x18\x02 \x01(\t\x12\x0f\n\x07\x63\x61pital\x18\x03 \x01(\t\x12\x11\n\tsubregion\x18\x04 \x01(\t\x12\x12\n\npopulation\x18\x05 \x01(\x05\x12\x12\n\nnativeName\x18\x06 \x01(\t\x12)\n\ncurrencies\x18\x07 \x03(\x0b\x32\x15.countries.Currencies2L\n\x07\x43ountry\x12\x41\n\x06Search\x12\x19.countries.CountryRequest\x1a\x1a.countries.CountryResponse\"\x00\x62\x06proto3')



_COUNTRYREQUEST = DESCRIPTOR.message_types_by_name['CountryRequest']
_CURRENCIES = DESCRIPTOR.message_types_by_name['Currencies']
_COUNTRYRESPONSE = DESCRIPTOR.message_types_by_name['CountryResponse']
CountryRequest = _reflection.GeneratedProtocolMessageType('CountryRequest', (_message.Message,), {
  'DESCRIPTOR' : _COUNTRYREQUEST,
  '__module__' : 'countries_pb2'
  # @@protoc_insertion_point(class_scope:countries.CountryRequest)
  })
_sym_db.RegisterMessage(CountryRequest)

Currencies = _reflection.GeneratedProtocolMessageType('Currencies', (_message.Message,), {
  'DESCRIPTOR' : _CURRENCIES,
  '__module__' : 'countries_pb2'
  # @@protoc_insertion_point(class_scope:countries.Currencies)
  })
_sym_db.RegisterMessage(Currencies)

CountryResponse = _reflection.GeneratedProtocolMessageType('CountryResponse', (_message.Message,), {
  'DESCRIPTOR' : _COUNTRYRESPONSE,
  '__module__' : 'countries_pb2'
  # @@protoc_insertion_point(class_scope:countries.CountryResponse)
  })
_sym_db.RegisterMessage(CountryResponse)

_COUNTRY = DESCRIPTOR.services_by_name['Country']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _COUNTRYREQUEST._serialized_start=30
  _COUNTRYREQUEST._serialized_end=60
  _CURRENCIES._serialized_start=62
  _CURRENCIES._serialized_end=118
  _COUNTRYRESPONSE._serialized_start=121
  _COUNTRYRESPONSE._serialized_end=291
  _COUNTRY._serialized_start=293
  _COUNTRY._serialized_end=369
# @@protoc_insertion_point(module_scope)
