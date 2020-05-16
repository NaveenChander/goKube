USE [GoKube]
GO

/****** Object:  StoredProcedure [experian].[usp_ExperianRequest_add]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_ExperianRequest_add')
BEGIN
    DROP PROCEDURE [experian].[usp_ExperianRequest_add]
END
GO

/****** Object:  StoredProcedure [experian].[usp_ExperianRequest_add]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-05-15  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [experian].[usp_ExperianRequest_add]
    (
    @ExperianRequestID bigint = 0,
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
    IF (@ExperianRequestID = 0 OR @Request is null OR TRIM(@Request) = '')
	BEGIN
        RAISERROR('Invalid Input parameters', 16,1);
        RETURN
    END


    INSERT INTO [experian].[ExperianRequest]
        ([ExperianRequestID]
        ,[Request]
        ,[Response]
        ,[RequestDate]
        ,[ResponseDate])
    VALUES
        (@ExperianRequestID
           , @Request
           , null
           , GETUTCDATE()
           , null)
END