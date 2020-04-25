USE [GoKube]
GO

ALTER TABLE [customer].[CustomerServiceOfferings] DROP CONSTRAINT [FK_CustomerServiceOfferings_CustomerServiceOfferings]
GO

/****** Object:  Table [customer].[CustomerServiceOfferings]    Script Date: 4/22/2020 4:46:10 PM ******/
DROP TABLE [customer].[CustomerServiceOfferings]
GO

/****** Object:  Table [customer].[CustomerServiceOfferings]    Script Date: 4/22/2020 4:46:10 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [customer].[CustomerServiceOfferings]
(
    [CustomerID] [int] NOT NULL,
    [CustomerServiceOfferingID] [int] NOT NULL,
    [CustomerClientProductID] [int] NOT NULL,
    [StartDate] [datetime] NOT NULL,
    [EndDate] [datetime] NULL,
    [BillingCycle] [nvarchar](50) NOT NULL,
    CONSTRAINT [PK_CustomerServiceType] PRIMARY KEY CLUSTERED 
(
	[CustomerID] ASC,
	[CustomerServiceOfferingID] ASC,
	[CustomerClientProductID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

ALTER TABLE [customer].[CustomerServiceOfferings]  WITH CHECK ADD  CONSTRAINT [FK_CustomerServiceOfferings_CustomerServiceOfferings] FOREIGN KEY([CustomerID], [CustomerServiceOfferingID], [CustomerClientProductID])
REFERENCES [customer].[CustomerServiceOfferings] ([CustomerID], [CustomerServiceOfferingID], [CustomerClientProductID])
GO

ALTER TABLE [customer].[CustomerServiceOfferings] CHECK CONSTRAINT [FK_CustomerServiceOfferings_CustomerServiceOfferings]
GO

