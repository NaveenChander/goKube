USE [GoKube]
GO

/****** Object:  Table [Experian].[ExperianRequest]    Script Date: 5/15/2020 5:58:07 PM ******/
DROP TABLE [experian].[ExperianRequest]
GO

/****** Object:  Table [Experian].[ExperianRequest]    Script Date: 5/15/2020 5:58:07 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [experian].[ExperianRequest]
(
    [ExperianRequestID] [bigint] NOT NULL,
    [Request] [nvarchar](max) NOT NULL,
    [Response] [nvarchar](max) NULL,
    [RequestDate] [datetime] NOT NULL,
    [ResponseDate] [datetime] NULL,
    CONSTRAINT [PK_ExperianRequest] PRIMARY KEY CLUSTERED 
(
	[ExperianRequestID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

