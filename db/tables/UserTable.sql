USE [GoKube]
GO

/****** Object:  Table [auth].[UserTable]    Script Date: 4/22/2020 5:48:03 PM ******/
DROP TABLE [auth].[UserTable]
GO

/****** Object:  Table [auth].[UserTable]    Script Date: 4/22/2020 5:48:03 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [auth].[UserTable]
(
    [UserName] [nvarchar](100) NOT NULL,
    [Permissions] [nvarchar](500) NOT NULL
) ON [PRIMARY]
GO

