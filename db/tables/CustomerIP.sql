USE [GoKube]
GO

ALTER TABLE [customer].[CustomerIP] DROP CONSTRAINT [FK_CustomerIP_Customer]
GO

/****** Object:  Table [customer].[CustomerIP]    Script Date: 4/22/2020 4:40:11 PM ******/
DROP TABLE [customer].[CustomerIP]
GO

/****** Object:  Table [customer].[CustomerIP]    Script Date: 4/22/2020 4:40:11 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [customer].[CustomerIP]
(
    [CustomerIP] [nvarchar](100) NOT NULL,
    [CustomerID] [int] NOT NULL,
    [StartDate] [datetime] NOT NULL,
    [EndDate] [datetime] NOT NULL,
    [RowInserted] [datetime] NOT NULL,
    CONSTRAINT [PK_CustomerID_CustomerIP] PRIMARY KEY CLUSTERED 
(
	[CustomerIP] ASC,
	[CustomerID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

ALTER TABLE [customer].[CustomerIP]  WITH CHECK ADD  CONSTRAINT [FK_CustomerIP_Customer] FOREIGN KEY([CustomerID])
REFERENCES [customer].[Customer] ([CustomerID])
GO

ALTER TABLE [customer].[CustomerIP] CHECK CONSTRAINT [FK_CustomerIP_Customer]
GO

