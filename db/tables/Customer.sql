USE [GoKube]
GO

/****** Object:  Table [customer].[Customer]    Script Date: 4/22/2020 4:38:22 PM ******/
DROP TABLE [customer].[Customer]
GO

/****** Object:  Table [customer].[Customer]    Script Date: 4/22/2020 4:38:22 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [customer].[Customer]
(
    [CustomerID] [int] NOT NULL,
    [Name] [nvarchar](200) NOT NULL,
    [InternalCustomerID] [nvarchar](50) NOT NULL,
    [StartDate] [datetime] NOT NULL,
    [EndDate] [datetime] NULL,
    [InsertedDate] [datetime] NOT NULL,
    [RowVersion] [timestamp] NOT NULL,
    CONSTRAINT [PK_Customer] PRIMARY KEY CLUSTERED 
(
	[CustomerID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

