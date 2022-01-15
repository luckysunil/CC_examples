import avro.schema
from avro.datafile import DataFileReader, DataFileWriter
from avro.io import DatumReader, DatumWriter

#https://avro.apache.org/docs/current/spec.html
schema = avro.schema.parse(open("sample_schema.avsc", "rb").read())

writer = DataFileWriter(open("sample_schema.avro", "wb"), DatumWriter(), schema)

sample_array_data = ["One String","Two String"]

writer.append({
                "sample_string":"Sunil",
                "sample_int":0,
                "sample_float":100.50,
                "sample_boolean":True,
                "sample_string_union":"Kumar",
                "sample_int_union":100,
                "sample_nested_record":{
                    "sample_nested_string":"Nested String Data",
                    "sample_nested_int":10
                    },
                "sample_array":sample_array_data
                })

writer.append({
                "sample_string":"Jasmine",
                "sample_int":1,
                "sample_float":1.5,
                "sample_boolean":False,
                "sample_nested_record":{
                    "sample_nested_string":"Nested String Data",
                    "sample_nested_int":10
                    },
                "sample_array":sample_array_data
                })

writer.close()

reader = DataFileReader(open("sample_schema.avro", "rb"), DatumReader())
for data in reader:
    print(data)

reader.close()

"""

# Sample Output
{'sample_string': 'Sunil', 'sample_int': 0, 'sample_float': 100.5, 'sample_boolean': True, 'sample_string_union': 'Kumar', 'sample_int_union': 100, 'sample_nested_record': {'sample_nested_string': 'Nested String Data', 'sample_nested_int': 10}, 'sample_array': ['One String', 'Two String']}

{'sample_string': 'Jasmine', 'sample_int': 1, 'sample_float': 1.5, 'sample_boolean': False, 'sample_string_union': None, 'sample_int_union': None, 'sample_nested_record': {'sample_nested_string': 'Nested String Data', 'sample_nested_int': 10}, 'sample_array': ['One String', 'Two String']}

"""
