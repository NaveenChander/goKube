USE [GoKube]
GO

/****** Object:  StoredProcedure [experian].[usp_ExperianRequest_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_ExperianRequest_sel')
BEGIN
    DROP PROCEDURE [experian].[usp_ExperianRequest_sel]
END
GO

/****** Object:  StoredProcedure [experian].[usp_ExperianRequest_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-05-15  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [experian].[usp_ExperianRequest_sel]
    (
    @ExperianRequestID bigint = 0,
    @FromRequestDate datetime = '01/01/1900',
    @ToRequestDate datetime = '01/01/1900'
)
AS
BEGIN

    SET NOCOUNT ON;

    SELECT [ExperianRequestID]
      , [Request]
      , [Response]
      , [RequestDate]
      , [ResponseDate]
    FROM [experian].[ExperianRequest]
    WHERE 
		( @ExperianRequestID = 0 OR ExperianRequestID = @ExperianRequestID)
        OR
        ( @ExperianRequestID = 0 AND @ToRequestDate = '01/01/1900' AND [RequestDate] = @FromRequestDate)
        OR
        (@ExperianRequestID = 0 AND '01/01/1900' NOT IN (@FromRequestDate, @ToRequestDate) AND [RequestDate] BETWEEN @FromRequestDate AND @ToRequestDate)

END

GO