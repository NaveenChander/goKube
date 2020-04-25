USE [GoKube]
GO

/****** Object:  Table [customer].[CustomerProduct]    Script Date: 4/22/2020 4:43:50 PM ******/
DROP TABLE [customer].[CustomerProduct]
GO

/****** Object:  Table [customer].[CustomerProduct]    Script Date: 4/22/2020 4:43:50 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [customer].[CustomerProduct]
(
    [CustomerProductID] [int] NOT NULL,
    [CustomerProduct] [nvarchar](100) NOT NULL,
    CONSTRAINT [PK_CustomerProduct] PRIMARY KEY CLUSTERED 
(
	[CustomerProductID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

