import avro.schema
from avro.datafile import DataFileReader, DataFileWriter
from avro.io import DatumReader, DatumWriter

schema = avro.schema.parse(open("sample_schema.avsc", "rb").read())

writer = DataFileWriter(open("sample_schema.avro", "wb"), DatumWriter(), schema)
writer.append({"sample_string":"Jack","sample_int_union":256})
writer.append({"sample_string":"Ben","sample_int":7,"sample_string_union":"red"})
writer.close()

reader = DataFileReader(open("sample_schema.avro", "rb"), DatumReader())
for data in reader:
    print(data)

reader.close()
