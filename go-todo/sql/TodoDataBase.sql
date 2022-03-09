USE [master]
GO
/****** Object:  Database [ToDoDB]    Script Date: 03/07/2022 09:38:06 ******/
CREATE DATABASE [ToDoDB]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'test', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL15.MSSQLSERVER01\MSSQL\DATA\test.mdf' , SIZE = 8192KB , MAXSIZE = UNLIMITED, FILEGROWTH = 65536KB )
 LOG ON 
( NAME = N'test_log', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL15.MSSQLSERVER01\MSSQL\DATA\test_log.ldf' , SIZE = 8192KB , MAXSIZE = 2048GB , FILEGROWTH = 65536KB )
 WITH CATALOG_COLLATION = DATABASE_DEFAULT
GO
ALTER DATABASE [ToDoDB] SET COMPATIBILITY_LEVEL = 150
GO
IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [ToDoDB].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO

USE [ToDoDB]
GO
/****** Object:  User [test]    Script Date: 03/07/2022 09:38:06 ******/
CREATE USER [test] FOR LOGIN [test] WITH DEFAULT_SCHEMA=[dbo]
GO
ALTER ROLE [db_datareader] ADD MEMBER [test]
GO
ALTER ROLE [db_datawriter] ADD MEMBER [test]
GO
/****** Object:  Schema [TestSchema]    Script Date: 03/07/2022 09:38:06 ******/
CREATE SCHEMA [TestSchema]
GO
/****** Object:  Table [dbo].[tasks]    Script Date: 03/07/2022 09:38:06 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[tasks](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[task] [nvarchar](50) NOT NULL,
	[status] [bit] NOT NULL,
	[createdDateTime] [datetime] NOT NULL,
	[lastUpdatedDateTime] [datetime] NOT NULL
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[tasks] ADD  CONSTRAINT [DF_tasks_status]  DEFAULT ((0)) FOR [status]
GO
/****** Object:  StoredProcedure [dbo].[deleteOneTask]    Script Date: 03/07/2022 09:38:06 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[deleteOneTask]
(@id bigint)
AS
BEGIN
DELETE FROM tasks WHERE id = @id
END
GO
/****** Object:  StoredProcedure [dbo].[getAllTask]    Script Date: 03/07/2022 09:38:06 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[getAllTask]
AS
BEGIN
SELECT [id]
      ,[task]
      ,[status]
      ,[createdDateTime]
      ,[lastUpdatedDateTime]
  FROM [dbo].[tasks]
  order by createdDateTime asc
END
GO
/****** Object:  StoredProcedure [dbo].[InsertOneTask]    Script Date: 03/07/2022 09:38:06 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[InsertOneTask]
(@task nvarchar(50),
 @id bigint OUTPUT
)
AS
BEGIN
	INSERT INTO tasks(
	task,
	[status],
	createdDateTime,
	lastUpdatedDateTime)

	VALUES(
	@task,
	0,
	GETDATE(),
	GETDATE())

   select ID = convert(bigint, SCOPE_IDENTITY())

end
GO
/****** Object:  StoredProcedure [dbo].[updateTask]    Script Date: 03/07/2022 09:38:06 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE   PROCEDURE [dbo].[updateTask]
(
 @id bigint,
 @status bit,
 @task nvarchar(50)
 )
AS
BEGIN
	UPDATE tasks
	SET	[status] = @status,
		task = ISNULL(NULLIF(LTRIM(RTRIM(@task)), ''), task),
		lastUpdatedDateTime = GETDATE()
	WHERE 
		id = @id
END
GO
USE [master]
GO
ALTER DATABASE [ToDoDB] SET  READ_WRITE 
GO
