USE [GoKube]
GO

/****** Object:  Table [idMatch].[IDMatchRequest]    Script Date: 5/15/2020 5:58:07 PM ******/
DROP TABLE [idMatch].[IDMatchRequest]
GO

/****** Object:  Table [idMatch].[IDMatchRequest]    Script Date: 5/15/2020 5:58:07 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [idMatch].[IDMatchRequest]
(
    [IDMatchRequestID] [bigint] NOT NULL,
    [Request] [nvarchar](max) NOT NULL,
    [Response] [nvarchar](max) NULL,
    [RequestDate] [datetime] NOT NULL,
    [ResponseDate] [datetime] NULL,
    CONSTRAINT [PK_IDMatchRequest] PRIMARY KEY CLUSTERED 
(
	[IDMatchRequestID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

