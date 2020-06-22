USE [GoKube]
GO

/****** Object:  StoredProcedure [IDMatch].[usp_IDMatchRequest_upd]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_IDMatchRequest_upd')
BEGIN
    DROP PROCEDURE [idMatch].[usp_IDMatchRequest_upd]
END
GO

/****** Object:  StoredProcedure [idMatch].[usp_IDMatchRequest_upd]    Script Date: 05/15/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-06-22  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [idMatch].[usp_IDMatchRequest_upd]
    (
    @IDMatchRequestID nvarchar(100),
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
    FROM [IDMatch].IDMatchRequest
    WHERE IDMatchRequestID = @IDMatchRequestID)
	BEGIN

        UPDATE 
		[IDMatch].[IDMatchRequest]
	SET 
		[Response] = @Response
		,[ResponseDate] = GETUTCDATE()
	WHERE 
		[IDMatchRequestID] = @IDMatchRequestID
    END
	ELSE
	BEGIN
        RAISERROR('IDMatchRequestID does not exists',16,1);
        RETURN;
    END

END

GO