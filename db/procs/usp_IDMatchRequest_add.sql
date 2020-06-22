USE [GoKube]
GO

/****** Object:  StoredProcedure [idMatch].[usp_IDMatchRequest_add]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_IDMatchRequest_add')
BEGIN
    DROP PROCEDURE [idMatch].[usp_IDMatchRequest_add]
END
GO

/****** Object:  StoredProcedure [idMatch].[usp_IDMatchRequest_add]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-06-22  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [idMatch].[usp_IDMatchRequest_add]
    (
    @IDMatchRequestID bigint = 0,
    @Request nvarchar(max)
)
AS
BEGIN

    SET NOCOUNT ON;

    IF @@TRANCOUNT = 0
	BEGIN
        RAISERROR('Procedure must be called within a transaction', 16,1);
        RETURN;
    END
    IF (@IDMatchRequestID = 0 OR @Request is null OR TRIM(@Request) = '')
	BEGIN
        RAISERROR('Invalid Input parameters', 16,1);
        RETURN
    END


    INSERT INTO [idMatch].[IDMatchRequest]
        ([IDMatchRequestID]
        ,[Request]
        ,[Response]
        ,[RequestDate]
        ,[ResponseDate])
    VALUES
        (@IDMatchRequestID
           , @Request
           , null
           , GETUTCDATE()
           , null)
END