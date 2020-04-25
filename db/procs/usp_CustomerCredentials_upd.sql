USE [GoKube]
GO

/****** Object:  StoredProcedure [Customer].[usp_CustomerCredentials_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_CustomerCredentials_upd')
BEGIN
    DROP PROCEDURE [Customer].[usp_CustomerCredentials_upd]
END
GO

/****** Object:  StoredProcedure [Customer].[usp_CustomerCredentials_upd]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-04-23  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [Customer].[usp_CustomerCredentials_upd]
    (
    @ClientAPIKey nvarchar(100),
    @Secret nvarchar(100) NULL,
    @startDate datetime NULL,
    @endDate datetime NULL,
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

    IF @ClientAPIKey IS NULL
	BEGIN
        RAISERROR('Client API cannot be NULL', 16,1);
        RETURN;
    END

    IF EXISTS (SELECT 1
    FROM [Customer].[CustomerCredentials]
    WHERE ClientAPIKey = @ClientAPIKey)
	BEGIN
        SELECT
            @Secret = ISNULL(@Secret, [Secret]),
            @startDate = ISNULL(@startDate, StartDate),
            @endDate = ISNULL(@endDate, EndDate)
        FROM
            CustomerCredentials
        WHERE 
			ClientAPIKey = @ClientAPIKey

        UPDATE
			CustomerCredentials 
		SET
			[Secret] = @Secret,
			StartDate = @StartDate,
			EndDate = @endDate
		WHERE
			ClientAPIKey = @ClientAPIKey

    END
	ELSE
	BEGIN
        RAISERROR('CustomerID does not exists',16,1);
        RETURN;
    END

END

GO