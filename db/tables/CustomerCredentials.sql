USE [GoKube]
GO

ALTER TABLE [customer].[CustomerCredentials] DROP CONSTRAINT [FK_CustomerCredentials_CustomerID_Customer_ID]
GO

/****** Object:  Table [customer].[CustomerCredentials]    Script Date: 4/22/2020 4:37:24 PM ******/
DROP TABLE [customer].[CustomerCredentials]
GO

/****** Object:  Table [customer].[CustomerCredentials]    Script Date: 4/22/2020 4:37:24 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [customer].[CustomerCredentials]
(
    [ClientAPIKey] [nvarchar](50) NOT NULL,
    [Secret] [nvarchar](200) NOT NULL,
    [CustomerID] [int] NOT NULL,
    [StartDate] [datetime] NOT NULL,
    [EndDate] [datetime] NULL,
    [RowInserted] [datetime] NOT NULL,
    CONSTRAINT [PK_CustomerCredentials_APIKey] PRIMARY KEY CLUSTERED 
(
	[ClientAPIKey] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

ALTER TABLE [customer].[CustomerCredentials]  WITH CHECK ADD  CONSTRAINT [FK_CustomerCredentials_CustomerID_Customer_ID] FOREIGN KEY([CustomerID])
REFERENCES [customer].[Customer] ([CustomerID])
GO

ALTER TABLE [customer].[CustomerCredentials] CHECK CONSTRAINT [FK_CustomerCredentials_CustomerID_Customer_ID]
GO

