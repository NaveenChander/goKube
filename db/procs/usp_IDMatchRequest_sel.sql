USE [GoKube]
GO

/****** Object:  StoredProcedure [idMatch].[usp_IDMatchRequest_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_IDMatchRequest_sel')
BEGIN
    DROP PROCEDURE [idMatch].[usp_IDMatchRequest_sel]
END
GO

/****** Object:  StoredProcedure [idMatch].[usp_IDMatchRequest_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-06-22  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [idMatch].[usp_IDMatchRequest_sel]
    (
    @IDMatchRequestID bigint = 0,
    @FromRequestDate datetime = '01/01/1900',
    @ToRequestDate datetime = '01/01/1900'
)
AS
BEGIN

    SET NOCOUNT ON;

    SELECT [IDMatchRequestID]
      , [Request]
      , [Response]
      , [RequestDate]
      , [ResponseDate]
    FROM [idMatch].[IDMatchRequest]
    WHERE 
		( @IDMatchRequestID = 0 OR IDMatchRequestID = @IDMatchRequestID)
        OR
        ( @IDMatchRequestID = 0 AND @ToRequestDate = '01/01/1900' AND [RequestDate] = @FromRequestDate)
        OR
        (@IDMatchRequestID = 0 AND '01/01/1900' NOT IN (@FromRequestDate, @ToRequestDate) AND [RequestDate] BETWEEN @FromRequestDate AND @ToRequestDate)

END

GO