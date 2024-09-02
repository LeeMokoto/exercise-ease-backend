resource "aws_rds_cluster" "dala" {
  cluster_identifier = "exercise-ease-cluster"
  engine             = "aurora-postgresql"
  engine_mode        = "provisioned"
  engine_version     = "16.3"
  database_name      = "exercise_ease_db"
  master_username    = "postgres"
  master_password    = ",HTjgKRo"
  skip_final_snapshot = true
  storage_encrypted  = true
  vpc_security_group_ids = [aws_security_group.rds.id]
  db_subnet_group_name = aws_db_subnet_group.aurora_subnet_group.name
  

  serverlessv2_scaling_configuration {
    max_capacity = 3.0
    min_capacity = 0.5
  }
}

resource "aws_rds_cluster_instance" "dala" {
  cluster_identifier = aws_rds_cluster.dala.id
  instance_class     = "db.serverless"
  engine             = aws_rds_cluster.dala.engine
  engine_version     = aws_rds_cluster.dala.engine_version
  publicly_accessible = true
  db_subnet_group_name = aws_db_subnet_group.aurora_subnet_group.name
  
}

resource "aws_db_subnet_group" "aurora_subnet_group" {
  name       = "tf-rds-exercise-ease"
  subnet_ids = [for subnet in aws_subnet.public : subnet.id]
  

}

output "rds-endpoint" {
    value = aws_rds_cluster_instance.dala.endpoint
}

resource "aws_security_group" "rds" {
  vpc_id      = aws_vpc.this.id
  name        = "${var.project_name}-rds-sg"
  description = "Allow inbound access"
  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "${var.project_name}-rds-sg"
  }
}