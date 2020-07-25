# t2.micro node with an AWS Tag naming it "HelloWorld"
provider "aws" {
  region = "us-west-2"
}

## pega recursos existentes
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_key_pair" "deployer" {
  key_name   = "deploy-key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7L72+A1k7FbO3B+tqRCNdu5OmENH01djjgMx/soMI/Wsy7a/C9cQg5RsYBazzrBamL9/79XGZDE6YUQJbAFaeXtfBxsFg+b3TDc/VFPMhT75CTqui3C9ZD9rMMN4wIAhyFjo+orQWR9XTYW6Fvzz0qlGkFZe6UM7hV+GjPcAxlssArlKu0nQpxM2CU1ekT8iA5ooujCgpGmeI20MCFDBSITcriTOKa1XIPc/His0yVx3ArDJxXxtA/7YcWDKjG85w3lkiNsyyp9v+nSVOGtRaiNNS4HxHpXwjjJY7sxDUeDQzE3l8ZSwzKTx0i73obv+gv854K3tBF4LGKRqv6+Q9"
}

## criação de novos recursos
resource "aws_instance" "web" {
  ami           = "${data.aws_ami.ubuntu.id}"
  instance_type = "t2.micro"
  key_name      = aws_key_pair.deployer.key_name
  user_data     = << EOF
		#! / bin / bash
        sudo apt-get update -y
		sudo apt install docker.io -y
  EOF

  tags = {
    Name = "Workshop"
    Billing = "Dev"
  }
}
