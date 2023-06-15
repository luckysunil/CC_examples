#Setup VPC
resource "aws_vpc" "gse_vpc" {
  cidr_block = var.cidr_block
  tags = {
    Name = "GSE_vpc"
  }
}

#Create Subnets
resource "aws_subnet" "public_subnets" {
 count             = length(var.public_subnet_cidrs)
 vpc_id            = aws_vpc.gse_vpc.id
 cidr_block        = element(var.public_subnet_cidrs, count.index)
 availability_zone = element(var.azs, count.index)

 tags = {
   Name = "gse_public_${count.index + 1}"
 }
}

resource "aws_subnet" "private_subnets" {
 count             = length(var.private_subnet_cidrs)
 vpc_id            = aws_vpc.gse_vpc.id
 cidr_block        = element(var.private_subnet_cidrs, count.index)
 availability_zone = element(var.azs, count.index)

 tags = {
   Name = "gse_private_${count.index + 1}"
 }
}

resource "aws_internet_gateway" "gse_igw" {
 vpc_id = aws_vpc.gse_vpc.id

 tags = {
   Name = "GSE_igw"
 }
}

resource "aws_route_table" "gse_public_rtb" {
 vpc_id = aws_vpc.gse_vpc.id

 route {
   cidr_block = "0.0.0.0/0"
   gateway_id = aws_internet_gateway.gse_igw.id
 }

 tags = {
   Name = "GSE_Public_Rtb"
 }
}

resource "aws_route_table_association" "gse_pub_rtb_subnet_assoc" {
 count = length(var.public_subnet_cidrs)
 subnet_id      = element(aws_subnet.public_subnets[*].id, count.index)
 route_table_id = aws_route_table.gse_public_rtb.id
}
