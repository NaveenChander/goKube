USE [GoKube]
GO

/****** Object:  StoredProcedure [experian].[usp_CustomerCredentials_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_ExperianRequest_upd')
BEGIN
    DROP PROCEDURE [experian].[usp_ExperianRequest_upd]
END
GO

/****** Object:  StoredProcedure [experian].[usp_ExperianRequest_upd]    Script Date: 05/15/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-05-15  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [experian].[usp_ExperianRequest_upd]
    (
    @experianRequestID nvarchar(100),
    @Response nvarchar(max),
    @UnitTestBypass bit=0
)
AS
BEGIN

    SET NOCOUNT ON;

    IF @@TRANCOUNT = 0
	BEGIN
        RAISERROR('Procedure must be called within a transaction', 16,1);
        RETURN;
    END

    IF @Response IS NULL OR TRIM(@Response) = ''
	BEGIN
        RAISERROR('Empty Response', 16,1);
        RETURN;
    END

    IF EXISTS (SELECT 1
    FROM [experian].ExperianRequest
    WHERE ExperianRequestID = @experianRequestID)
	BEGIN

        UPDATE 
		[experian].[ExperianRequest]
	SET 
		[Response] = @Response
		,[ResponseDate] = GETUTCDATE()
	WHERE 
		[ExperianRequestID] = @experianRequestID
    END
	ELSE
	BEGIN
        RAISERROR('ExperianRequestID does not exists',16,1);
        RETURN;
    END

END

GO