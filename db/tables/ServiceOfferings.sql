USE [GoKube]
GO

/****** Object:  Table [serviceOfferings].[ServiceOfferings]    Script Date: 4/22/2020 5:47:18 PM ******/
DROP TABLE [serviceOfferings].[ServiceOfferings]
GO

/****** Object:  Table [serviceOfferings].[ServiceOfferings]    Script Date: 4/22/2020 5:47:18 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [serviceOfferings].[ServiceOfferings]
(
    [ServiceOfferingID] [int] NOT NULL,
    [Name] [nvarchar](100) NOT NULL,
    [ServiceDescription] [varchar](500) NOT NULL,
    [InternalServiceRecordID] [nvarchar](100) NULL,
    [StartDate] [datetime] NOT NULL,
    [EndDate] [datetime] NULL,
    CONSTRAINT [PK_ServiceType_ID] PRIMARY KEY CLUSTERED 
(
	[ServiceOfferingID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

